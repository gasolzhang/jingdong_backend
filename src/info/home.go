package info

type HomeBanner struct {
	Image string
	//可能是url连接，也可能是商品id，也可能啥也没有那就只是单纯的图片
	Action string
}

type HomeGuess struct {
	Image   string
	GoodsId string
	Price   string
}
