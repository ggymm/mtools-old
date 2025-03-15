package model

import (
	"gorm.io/gorm"
	"mtools-pro/backend/pkg/consts"
)

type SnippetCatalog struct {
	Id         int64  `gorm:"column:id;primaryKey;autoIncrement;not null;comment:主键"`
	Name       string `gorm:"column:name;type:varchar(255);not null;comment:名称"`
	Extras     string `gorm:"column:extras;type:varchar(255);comment:额外参数"`
	ParentId   int64  `gorm:"column:parent_id;type:bigint(20);not null;comment:父节点ID"`
	ParentPath string `gorm:"column:parent_path;type:varchar(255);not null;comment:父节点路径"`
	BaseModel  `gorm:"embedded"`
}

func (m *SnippetCatalog) InitData(db *gorm.DB) {
	// 判断是否有数据，如果有则不初始化
	if db.First(&SnippetCatalog{}).RowsAffected > 0 {
		return
	}

	// 初始化数据
	var model SnippetCatalog
	model.Id = consts.DefaultId
	model.Name = "default"
	model.ParentId = consts.RootId

	db.Create(&model)
}
