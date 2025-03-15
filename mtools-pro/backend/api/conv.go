package api

type NumberBaseReq struct {
	Input     string `json:"input"`
	InputBase int    `json:"inputBase"`
}

type NumberBaseResp struct {
	Binary      string `json:"binary"`
	Octal       string `json:"octal"`
	Decimal     string `json:"decimal"`
	Hexadecimal string `json:"hexadecimal"`
}
