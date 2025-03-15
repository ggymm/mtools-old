package api

type MagnetConfigReq struct {
	Preload       int    `json:"preload"`
	RuleConfig    string `json:"ruleConfig"`
	CacheResult   int    `json:"cacheResult"`
	CacheTimeout  int    `json:"cacheTimeout"`
	ContentFilter string `json:"contentFilter"`

	ProxyHost   string `json:"proxyHost"`
	ProxyPort   int    `json:"proxyPort"`
	ProxyUser   string `json:"proxyUser"`
	ProxyPass   string `json:"proxyPass"`
	EnableProxy int    `json:"enableProxy"`
}

type MagnetConfigResp struct {
	Preload       int    `json:"preload"`
	RuleConfig    string `json:"ruleConfig"`
	CacheResult   int    `json:"cacheResult"`
	CacheTimeout  int    `json:"cacheTimeout"`
	ContentFilter string `json:"contentFilter"`

	ProxyHost   string `json:"proxyHost"`
	ProxyPort   int    `json:"proxyPort"`
	ProxyUser   string `json:"proxyUser"`
	ProxyPass   string `json:"proxyPass"`
	EnableProxy int    `json:"enableProxy"`
}

type MagnetSearchReq struct {
	PageReq
	Site     string `json:"site"`
	Keywords string `json:"keywords"`
}
