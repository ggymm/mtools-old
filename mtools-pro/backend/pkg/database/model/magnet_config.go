package model

import (
	"gorm.io/gorm"
	"mtools-pro/backend/pkg/consts"
)

type MagnetConfig struct {
	Id            int64  `gorm:"column:id;primaryKey;not null;comment:'主键'"`
	Preload       int    `gorm:"column:preload;type:tinyint(1);comment:'是否开启预加载'"`
	RuleConfig    string `gorm:"column:rule_config;type:varchar(255);comment:'自定义规则配置'"`
	CacheResult   int    `gorm:"column:cache_result;type:tinyint(1);comment:'是否缓存搜索结果'"`
	CacheTimeout  int    `gorm:"column:cache_timeout;type:int(11);comment:'缓存超时时间'"`
	ContentFilter string `gorm:"column:content_filter;type:varchar(255);comment:'内容过滤配置'"`

	ProxyHost   string `gorm:"column:proxy_host;type:varchar(255);comment:'代理地址'"`
	ProxyPort   int    `gorm:"column:proxy_port;type:int(11);comment:'代理端口'"`
	ProxyUser   string `gorm:"column:proxy_user;type:varchar(255);comment:'代理账号'"`
	ProxyPass   string `gorm:"column:proxy_pass;type:varchar(255);comment:'代理密码'"`
	EnableProxy int    `gorm:"column:enable_proxy;type:tinyint(1);comment:'是否启用代理'"`
}

func (m *MagnetConfig) InitData(db *gorm.DB) {
	// 判断是否有数据，如果有则不初始化
	if db.First(&MagnetConfig{}).RowsAffected > 0 {
		return
	}

	// 初始化数据
	var model MagnetConfig
	model.Id = consts.DefaultId

	model.Preload = 0
	model.RuleConfig = ""
	model.CacheResult = 0
	model.CacheTimeout = 0
	model.ContentFilter = ""

	model.ProxyHost = "127.0.0.1"
	model.ProxyPort = 9910
	model.ProxyUser = ""
	model.ProxyPass = ""
	model.EnableProxy = 1

	db.Create(&model)
}

func (m *MagnetConfig) GetProxyStruct() *struct {
	Host     string
	Port     int
	Username string
	Password string
} {
	return &struct {
		Host     string
		Port     int
		Username string
		Password string
	}{
		Host:     m.ProxyHost,
		Port:     m.ProxyPort,
		Username: m.ProxyUser,
		Password: m.ProxyPass,
	}
}
