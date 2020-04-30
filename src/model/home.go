package model

type HomeBanner struct {
	BaseModel
	Image string `gorm:"type:varchar(2084);not null;url"`
	//可能是url连接，也可能是商品id，也可能啥也没有那就只是单纯的图片
	Action string `gorm:"type:varchar(2084)"`
}

type HomeGuess struct {
	BaseModel
	Image   string `gorm:"type:varchar(2084);not null;url"`
	GoodsId string `gorm:"type:varchar(36);not null"`
	Price   string `gorm:"type:varchar(36);not null"`
}
