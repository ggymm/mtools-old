package magnet

import (
	"bytes"
	"fmt"
	"net/url"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/antchfx/htmlquery"
	"github.com/go-resty/resty/v2"
	"github.com/goccy/go-json"
	"github.com/pkg/errors"
	"golang.org/x/net/html"

	"mtools-pro/backend/assets"
	"mtools-pro/backend/pkg/log"
)

type Option struct {
	Page     int
	Site     string
	Keywords string

	Proxy *struct {
		Host     string
		Port     int
		Username string
		Password string
	}
}

type Info struct {
	Name         string `json:"name"`
	Size         string `json:"size"`
	Magnet       string `json:"magnet"`
	DownloadHot  string `json:"downloadHot"`
	DownloadTime string `json:"downloadTime"`
}

type Rule struct {
	Id      string `json:"id"`
	Name    string `json:"name"`
	Desc    string `json:"desc"`
	Proxy   bool   `json:"proxy"`
	Index   string `json:"index"`
	Referer string `json:"referer"`
	Publish string `json:"publish"`

	Path struct {
		Params  []PathParam `json:"params"`
		Default string      `json:"default"`
	} `json:"path"`

	Parse struct {
		Items        itemRule   `json:"items"`
		Count        []infoRule `json:"count"`
		Name         []infoRule `json:"name"`
		Size         []infoRule `json:"size"`
		Magnet       []infoRule `json:"magnet"`
		DownloadHot  []infoRule `json:"downloadHot"`
		DownloadTime []infoRule `json:"downloadTime"`
	}
}

type itemRule struct {
	// url xpath
	Xpath      string `json:"xpath"`
	StartIndex int    `json:"startIndex"`
}

type infoRule struct {
	// regex, xpath, xpathList, subIndex, subBetween
	Type    string `json:"type"`
	Regex   string `json:"regex"`
	Xpath   string `json:"xpath"`
	Start   int    `json:"start"`
	Finish  int    `json:"finish"`
	After   string `json:"after"`
	Before  string `json:"before"`
	Default string `json:"default"`
}

type PathParam struct {
	Func    string `json:"func"`
	Param   string `json:"param"`
	Default string `json:"default"`
}

func (info *Info) String() string {
	return fmt.Sprintf(
		"Name: %s, Size: %s, Magnet: %s, DownloadHot: %s, DownloadTime: %s",
		info.Name,
		info.Size,
		info.Magnet,
		info.DownloadHot,
		info.DownloadTime,
	)
}

func buildUrl(opt Option, rule Rule) string {
	search := url.QueryEscape(opt.Keywords)

	params := map[string]string{
		"search": search,
		"page":   fmt.Sprintf("%d", opt.Page),
	}

	// 判断是否有其他参数
	if len(rule.Path.Params) > 0 {
		for _, param := range rule.Path.Params {
			if len(param.Func) > 0 {
				params[param.Param] = runFunc(param.Func, rule, param)
			}
		}
	}

	return rule.Index + os.Expand(rule.Path.Default, func(key string) string {
		return params[key]
	})
}

func buildClient(url string, opt Option, rule Rule) *resty.Client {
	client := resty.New()

	var (
		proxy   string
		headers = make(map[string]string)
	)

	// referer
	if len(rule.Referer) == 0 {
		headers["Referer"] = url
	} else {
		headers["Referer"] = rule.Referer
	}

	headers["User-Agent"] = randomUserAgent()

	// 设置代理
	if opt.Proxy != nil {
		proxy = fmt.Sprintf(
			"http://%s:%d",
			opt.Proxy.Host,
			opt.Proxy.Port,
		)
	}

	client.SetProxy(proxy)
	client.SetHeaders(headers)
	client.SetTimeout(30 * time.Second)
	return client
}

func extractInfo(item *html.Node, rules []infoRule) string {
	content := ""
	// regex xpath xpathList subIndex
	for _, rule := range rules {
		switch rule.Type {
		case "regex":
			reg, err := regexp.Compile(rule.Regex)
			if err != nil {
				log.Error().
					Err(errors.WithStack(err)).
					Msg("regex compile failed")
				break
			}
			match := reg.FindStringSubmatch(htmlquery.InnerText(item))
			if len(match) > 1 {
				content = match[1]
			}
		case "xpath":
			node := htmlquery.FindOne(item, rule.Xpath)
			if node != nil {
				content = htmlquery.InnerText(node)
			}
		case "xpathList":
			nodes := htmlquery.Find(item, rule.Xpath)
			for _, node := range nodes {
				content += htmlquery.InnerText(node)
			}
		case "subIndex":
			start := rule.Start
			finish := rule.Finish
			if start == 0 && finish == 0 {
				break
			}

			length := len(content)
			if start < 0 {
				start += length
				start = max(start, 0)
			}
			if finish < 0 {
				finish += length
			}

			finish = max(finish, start)
			finish = min(finish, length)
			content = content[start:finish]
		case "subBetween":
			start := strings.Index(content, rule.After)
			finish := strings.Index(content, rule.Before)
			if start < 0 || finish < 0 {
				break
			}

			start = start + len(rule.After)
			finish = max(start, finish)
			content = content[start:finish]
		}
		if len(content) == 0 {
			content = rule.Default
		}
	}

	return content
}

func Search(opts ...Option) ([]Info, string, error) {
	var (
		err error
		opt Option

		rule  Rule
		rules = make([]Rule, 0)

		req  string
		resp *resty.Response

		doc   *html.Node
		items []*html.Node

		list  = make([]Info, 0)
		count string
	)

	defer func() {
		if r := recover(); r != nil {
			log.Error().
				Err(errors.WithStack(r.(error))).
				Msg("magnet search panic")
			return
		}
	}()

	if len(opts) <= 0 {
		return list, count, errors.New(OptIsEmpty)
	}
	opt = opts[0]

	err = json.Unmarshal([]byte(assets.MagnetRule), &rules)
	if err != nil {
		return list, count, err
	}

	// 1. 从 MagnetRule 中获取到对应的网站
	for _, r := range rules {
		if r.Id == opt.Site {
			rule = r
		}
	}
	if rule.Id == "" {
		return list, count, errors.New(SiteNotFound)
	}

	// 2. 拼接页面的访问地址
	req = buildUrl(opt, rule)
	log.Info().
		Str("url", req).
		Msg("magnet search url")

	// 3. 获取页面的文本内容
	resp, err = buildClient(req, opt, rule).R().Get(req)
	if err != nil {
		return list, count, err
	} else if resp.IsError() {
		return list, count, errors.New(RespStatusFailed)
	}

	// 4. 解析页面的文本内容
	doc, err = htmlquery.Parse(bytes.NewBuffer(resp.Body()))
	if err != nil {
		return list, count, err
	}

	count = extractInfo(doc, rule.Parse.Count)
	items = htmlquery.Find(doc, rule.Parse.Items.Xpath)

	log.Info().
		Str("count", count).
		Int("length", len(items)).
		Msg("magnet search result")
	for i := rule.Parse.Items.StartIndex; i < len(items); i++ {
		item := items[i]

		name := extractInfo(item, rule.Parse.Name)
		size := extractInfo(item, rule.Parse.Size)
		magnet := extractInfo(item, rule.Parse.Magnet)
		if len(magnet) > 0 {
			reg := regexp.MustCompile("^(/|.*.html)$")
			header := "magnet:?xt=urn:btih:"

			if reg.MatchString(magnet) {
				magnet = rule.Index + magnet
			} else if !strings.HasPrefix(magnet, header) {
				magnet = header + magnet
			}
		}
		downloadHot := extractInfo(item, rule.Parse.DownloadHot)
		downloadTime := extractInfo(item, rule.Parse.DownloadTime)

		log.Info().
			Str("name", name).
			Str("size", size).
			Str("magnet", magnet).
			Str("downloadHot", downloadHot).
			Str("downloadTime", downloadTime).
			Msg("magnet search result item")
		list = append(list, Info{
			Name:         name,
			Size:         size,
			Magnet:       magnet,
			DownloadHot:  downloadHot,
			DownloadTime: downloadTime,
		})
	}
	return list, count, nil
}
