package model

type MagnetCache struct {
	Id       int    `gorm:"column:id;type:int(11);not null;primaryKey;autoIncrement;comment:'主键'"`
	Keywords string `gorm:"column:keywords;type:varchar(255);not null;default:'';comment:'关键词'"`
}
