package model

type SnippetFragment struct {
	Id        int64  `gorm:"column:id;primaryKey;autoIncrement;not null;comment:'主键'"`
	SnippetId int64  `gorm:"column:snippet_id;type:bigint(20);not null;comment:'代码片段ID'"`
	Title     string `gorm:"column:title;type:varchar(255);not null;comment:'片段名称'"`
	Content   string `gorm:"column:content;type:longtext;not null;comment:'片段内容'"`
	Language  string `gorm:"column:language;type:varchar(255);not null;comment:'语言'"`
	BaseModel `gorm:"embedded"`
}
