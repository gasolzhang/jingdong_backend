package service

import (
	"errors"
	"github.com/gasolzhang/jdbackend/src/dao"
	"github.com/gasolzhang/jdbackend/src/model"
)

type GoodsService struct {
	dao *dao.GoodsDao
}

func NewGoodsService() *GoodsService {
	return &GoodsService{dao: dao.NewGoodsDao()}
}

func (this *GoodsService) InsertGoods(goods model.Goods) error {
	err := this.dao.InsertGoods(goods)
	if err != nil {
		return errors.New("插入商品失败")
	}
	return nil
}
