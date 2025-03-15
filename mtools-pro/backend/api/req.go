package api

type PageReq struct {
	Page int `query:"page"`
	Size int `query:"size"`
}

func (p *PageReq) GetPage() int {
	if p.Page <= 0 {
		p.Page = 1
	}
	return p.Page
}

func (p *PageReq) GetSize() int {
	if p.Size <= 0 {
		p.Size = 10
	}
	return p.Size
}

func (p *PageReq) GetOffset() int {
	return (p.GetPage() - 1) * p.GetSize()
}

type DeleteReq struct {
	Id int64 `json:"id" validate:"required"`
}
