package api

type DecodeQrcodeReq struct {
}

type DecodeQrcodeResp struct {
	Base64  string `json:"base64"`
	Content string `json:"content"`
}

type EncodeQrcodeReq struct {
	Size    int    `json:"size"`
	Content string `json:"content"`
}

type EncodeQrcodeResp struct {
	Base64 string `json:"base64"`
}
