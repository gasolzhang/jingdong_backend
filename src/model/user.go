package model

type User struct {
	BaseModel
	Phone string `gorm:"type:char(11);not null;unique"`
	//Age  sql.NullInt64 //零值类型
	Password string `gorm:"not null"`
	Name     string `gorm:"not null;type:varchar(20)"`
	Avatar   string `gorm:"type:varchar(2084)"`
	Sign     string `gorm:"type:varchar(30);nullable"`
}
