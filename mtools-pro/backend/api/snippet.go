package api

type SnippetCatalogTree struct {
	TreeNode[SnippetCatalogTree]
	Extras string `json:"extras"`
}

type SnippetCatalogCreateReq struct {
	Name     string `json:"name"`
	ParentId int64  `json:"parentId"`
}

type SnippetPageReq struct {
	PageReq
	CatalogId int64 `json:"catalogId"`
}

type SnippetPageResp struct {
	Id          int64  `json:"id"`
	Title       string `json:"title"`
	UpdateAt    string `json:"updateAt"`
	CatalogId   int64  `json:"catalogId"`
	CatalogName string `json:"catalogName"`
}

type SnippetReq struct {
	Id int64 `json:"id"`
}

type SnippetResp struct {
	Id        int64             `json:"id"`
	Title     string            `json:"title"`
	Fragments []SnippetFragment `json:"fragments"`
}

type SnippetCreateReq struct {
	Title     string            `json:"title"`
	CatalogId int64             `json:"catalogId"`
	Fragments []SnippetFragment `json:"fragments"`
}

type SnippetUpdateReq struct {
	Id        int64             `json:"id"`
	Title     string            `json:"title"`
	Fragments []SnippetFragment `json:"fragments"`
}

type SnippetFragment struct {
	Title    string `json:"title"`
	Content  string `json:"content"`
	Language string `json:"language"`
}
