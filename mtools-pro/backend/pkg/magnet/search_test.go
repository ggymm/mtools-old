package magnet

import (
	"regexp"
	"strings"
	"testing"

	"mtools-pro/backend/pkg/config"
	"mtools-pro/backend/pkg/log"
)

func Test_Search(t *testing.T) {
	//site := "btsow"
	//site := "bitsearch"
	site := "u3c3"
	//site := "yhg007"
	//site := "yuhuage"

	config.Init()
	log.Init()

	ops := Option{
		Page:     1,
		Site:     site,
		Keywords: "SSIS",

		Proxy: &struct {
			Host     string
			Port     int
			Username string
			Password string
		}{
			Host:     "127.0.0.1",
			Port:     9910,
			Username: "",
			Password: "",
		},
	}

	list, count, err := Search(ops)
	if err != nil {
		t.Fatal(err)
	}

	str := ""
	for _, item := range list {
		str += item.String() + "\n"
	}
	t.Log(str)
	t.Log(count)
}

func Test_Regex(t *testing.T) {
	reg := regexp.MustCompile("totalPages:\\s*(\\d+)")
	t.Log(reg.FindStringSubmatch("totalPages: 1111,"))
}

func Test_String(t *testing.T) {
	content := "//btsow.motorcycles/magnet/detail/hash/8B7C8B81D4F2CFA34AA9F7A7A7446F2E77CBBD7E"
	start := len(content) - 40
	finish := len(content)
	t.Log(content[start:finish])

	content = "/search/肖申克的救赎-7.html"
	start = strings.LastIndex(content, "-")
	finish = strings.LastIndex(content, ".")
	t.Log(content[start+1 : finish])
}
