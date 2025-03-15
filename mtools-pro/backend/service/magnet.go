package service

import (
	"context"
	"sync"

	"github.com/ggymm/gopkg/conv"
	"github.com/ggymm/gopkg/sys/files"
	"github.com/goccy/go-json"
	"github.com/wailsapp/wails/v2/pkg/runtime"

	"mtools-pro/backend/api"
	"mtools-pro/backend/assets"
	"mtools-pro/backend/pkg/consts"
	"mtools-pro/backend/pkg/database"
	"mtools-pro/backend/pkg/database/model"
	"mtools-pro/backend/pkg/magnet"
)

type MagnetService struct {
	ctx context.Context
}

var magnetOnce sync.Once
var magnetService *MagnetService

func NewMagnetService() *MagnetService {
	if magnetService == nil {
		magnetOnce.Do(func() {
			magnetService = &MagnetService{}
		})
	}
	return magnetService
}

func (s *MagnetService) Startup(ctx context.Context) {
	s.ctx = ctx
}

func (s *MagnetService) ChooseRuleConfig() api.Resp {
	fp, err := runtime.OpenFileDialog(s.ctx, runtime.OpenDialogOptions{
		Title: "选择自定义规则",
		Filters: []runtime.FileFilter{
			{
				Pattern:     "*.json",
				DisplayName: "JSON Files (*.json)",
			},
		},
		ShowHiddenFiles: true,
	})
	if err != nil {
		return api.Error(err.Error())
	}
	return api.Success(fp)
}

func (s *MagnetService) ChooseContentFilterConfig() api.Resp {
	fp, err := runtime.OpenFileDialog(s.ctx, runtime.OpenDialogOptions{
		Title:           "选择图片",
		ShowHiddenFiles: true,
	})
	if err != nil {
		return api.Error(err.Error())
	}
	return api.Success(fp)
}

func (s *MagnetService) GetConfig() api.Resp {
	config := model.MagnetConfig{}
	err := database.DB.First(&config, consts.DefaultId).Error
	if err != nil {
		return api.Error(err.Error())
	}

	// 构建响应值
	resp := api.MagnetConfigResp{
		Preload:       config.Preload,
		RuleConfig:    config.RuleConfig,
		CacheResult:   config.CacheResult,
		CacheTimeout:  config.CacheTimeout,
		ContentFilter: config.ContentFilter,
		ProxyHost:     config.ProxyHost,
		ProxyPort:     config.ProxyPort,
		ProxyUser:     config.ProxyUser,
		ProxyPass:     config.ProxyPass,
		EnableProxy:   config.EnableProxy,
	}
	return api.Success(resp)
}

func (s *MagnetService) UpdateConfig(req api.MagnetConfigReq) api.Resp {
	config := model.MagnetConfig{
		Id:            consts.DefaultId,
		Preload:       req.Preload,
		RuleConfig:    req.RuleConfig,
		CacheResult:   req.CacheResult,
		CacheTimeout:  req.CacheTimeout,
		ContentFilter: req.ContentFilter,
		ProxyHost:     req.ProxyHost,
		ProxyPort:     req.ProxyPort,
		ProxyUser:     req.ProxyUser,
		ProxyPass:     req.ProxyPass,
		EnableProxy:   req.EnableProxy,
	}
	err := database.DB.Save(&config).Error
	if err != nil {
		return api.Error(err.Error())
	}
	return api.Success(nil)
}

func (s *MagnetService) GetRuleOptions() api.Resp {
	var (
		err      error
		config   model.MagnetConfig
		ruleJson string

		rules   = make([]magnet.Rule, 0)
		options = make([]api.OptionResp, 0)
	)
	err = database.DB.First(&config, consts.DefaultId).Error
	if err != nil {
		return api.Error(err.Error())
	}

	// 判断是否存在自定义规则
	if config.RuleConfig != "" {
		// 读取自定义规则
		ruleJson, err = files.ReadString(config.RuleConfig)
		if err != nil {
			return api.Error(err.Error())
		}
	} else {
		ruleJson = assets.MagnetRule
	}

	err = json.Unmarshal([]byte(ruleJson), &rules)
	if err != nil {
		return api.Error(err.Error())
	}

	for _, rule := range rules {
		options = append(options, api.OptionResp{
			Value: rule.Id,
			Label: rule.Name,
		})
	}
	return api.Success(options)
}

func (s *MagnetService) RunSearchMagnet(req api.MagnetSearchReq) api.Resp {
	var (
		err    error
		config model.MagnetConfig

		list  = make([]magnet.Info, 0)
		total string
	)

	// 读取配置
	err = database.DB.First(&config, consts.DefaultId).Error
	if err != nil {
		return api.Error(err.Error())
	}

	ops := magnet.Option{
		Page:     req.Page,
		Site:     req.Site,
		Keywords: req.Keywords,
		Proxy:    config.GetProxyStruct(),
	}
	list, total, err = magnet.Search(ops)
	if err != nil {
		return api.Error(err.Error())
	}
	go func() {
		// 缓存数据
		if config.CacheResult == consts.Enable {
		}
	}()
	return api.Success(api.PageResp{List: list, Total: conv.ToInt64(total)})
}
