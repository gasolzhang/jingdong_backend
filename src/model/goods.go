package model

type Goods struct {
	BaseModel
	Cover     string `gorm:"type:varchar(2084);not null;url"`
	Name      string `gorm:"type:varchar(1024);not null"`
	Price     string `gorm:"type:varchar(10);not null"`
	OldPrice  string `gorm:"type:varchar(10);not null"`
	CateId    string `gorm:"type:varchar(36);not null"`
	SaleCount int
}
