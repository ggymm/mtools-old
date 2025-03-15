package database

import (
	_ "embed"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"

	"mtools-pro/backend/pkg/database/model"
)

var (
	DB *gorm.DB
)

//go:embed db.cfg
var datasource string

func Init() (err error) {
	DB, err = gorm.Open(mysql.Open(datasource), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		Logger: NewCustomLog().LogMode(logger.Info),

		DisableForeignKeyConstraintWhenMigrating: true, // https://gorm.io/zh_CN/docs/constraints.html
	})
	if err != nil {
		return err
	}

	err = DB.AutoMigrate(
		&model.MagnetCache{},
		&model.MagnetConfig{},

		&model.Snippet{},
		&model.SnippetCatalog{},
		&model.SnippetFragment{},
	)
	if err != nil {
		return err
	}

	initTable(&model.MagnetConfig{})
	initTable(&model.SnippetCatalog{})
	return
}

func initTable(table interface{}) {
	if DB.Migrator().HasTable(table) {
		// 判断是否有InitData方法，如果有则执行
		if fn, ok := table.(interface{ InitData(*gorm.DB) }); ok {
			fn.InitData(DB)
		}
	}
}
