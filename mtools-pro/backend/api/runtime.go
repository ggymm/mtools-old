package api

type FileFilter struct {
	Pattern     string `json:"pattern"`
	DisplayName string `json:"displayName"`
}

type ChooseFileReq struct {
	Title  string       `json:"title"`
	Filter []FileFilter `json:"filter"`
}
