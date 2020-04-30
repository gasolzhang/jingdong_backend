package service

import (
	"errors"
	"github.com/gasolzhang/jdbackend/src/dao"
	"github.com/gasolzhang/jdbackend/src/info"
	"github.com/gasolzhang/jdbackend/src/model"
)

type HomeService struct {
	dao *dao.HomeDao
}

func NewHomeService() *HomeService {
	return &HomeService{dao: dao.NewHomeDao()}
}

func (this *HomeService) GetBanner(limit int) (banner []info.HomeBanner, err error) {
	banner, err = this.dao.GetBanner(limit)
	if err != nil {
		return nil, errors.New("查询首页Banner失败")
	}
	return banner, nil
}

func (this *HomeService) InsertBanner(image, action string) error {
	return this.dao.InsertBanner(&model.HomeBanner{Image: image, Action: action})
}

func (this *HomeService) InsertGuess(image, goodsId, price string) error {
	return this.dao.InsertGuess(&model.HomeGuess{Image: image, GoodsId: goodsId, Price: price})
}

func (this *HomeService) GetGuess(limit int) (guess []info.HomeGuess, err error) {
	guess, err = this.dao.GetGuess(limit)
	if err != nil {
		return nil, errors.New("查询首页Guess失败")
	}
	return guess, nil
}

func (this *HomeService) GetHotGoods(page, limit int) (goods []info.Goods, err error) {
	goods, err = this.dao.GetHotGoods(page, limit)
	if err != nil {
		return nil, errors.New("查询热门商品失败")
	}
	return goods, nil
}
