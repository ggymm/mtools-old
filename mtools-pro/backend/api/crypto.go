package api

type BcryptEncodeReq struct {
	Input     string `json:"input"`
	SaltCount int    `json:"saltCount"`
}

type BcryptCompareReq struct {
	Origin string `json:"origin"`
	Encode string `json:"encode"`
}
