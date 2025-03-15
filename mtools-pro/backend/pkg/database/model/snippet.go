package model

import "time"

type Snippet struct {
	Id        int64  `gorm:"column:id;primaryKey;autoIncrement;not null;comment:'主键'"`
	Title     string `gorm:"column:title;type:varchar(255);not null;comment:'名称'"`
	CatalogId int64  `gorm:"column:catalog_id;type:bigint(20);not null;comment:'目录ID'"`
	BaseModel `gorm:"embedded"`
}

func (h *Snippet) GetUpdateAt() string {
	return time.UnixMilli(h.UpdatedAt).Format("2006-01-02 15:04:05")
}

type SnippetExt struct {
	Snippet
	SnippetCatalog SnippetCatalog `gorm:"foreignKey:CatalogId"`
}

func (SnippetExt) TableName() string {
	return "snippet"
}
