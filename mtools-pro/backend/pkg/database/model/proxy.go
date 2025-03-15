package model

type Proxy struct {
	Id        int64  `gorm:"column:id;primaryKey;autoIncrement:true;not null;comment:ID"`
	Type      string `gorm:"column:type;type:varchar(32);not null;comment:类型"`
	Host      string `gorm:"column:host;type:varchar(32);not null;comment:主机"`
	Port      int    `gorm:"column:port;type:int(11);not null;comment:端口"`
	Username  string `gorm:"column:username;type:varchar(32);not null;comment:用户名"`
	Password  string `gorm:"column:password;type:varchar(32);not null;comment:密码"`
	BaseModel `gorm:"embedded"`
}
