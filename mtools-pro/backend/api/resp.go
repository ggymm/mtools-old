package api

type TreeNode[T any] struct {
	Key      int64  `json:"key"`
	Label    string `json:"label"`
	Children []*T   `json:"children"`
}

type PageResp struct {
	List  interface{} `json:"list"`
	Total int64       `json:"total"`
}

type OptionResp struct {
	Value string `json:"value"`
	Label string `json:"label"`
}

type Resp struct {
	Msg     string `json:"msg,omitempty"`
	Data    any    `json:"data,omitempty"`
	Success bool   `json:"success"`
}

func Error(msg string) (r Resp) {
	r.Msg = msg
	r.Success = false
	return r
}

func Success(data any) (r Resp) {
	r.Data = data
	r.Success = true
	return r
}
