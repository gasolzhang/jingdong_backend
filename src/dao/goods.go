package dao

import (
	"github.com/gasolzhang/jdbackend/src/model"
	"github.com/gasolzhang/jdbackend/src/mygorm"
	"github.com/jinzhu/gorm"
)

type GoodsDao struct {
	db *gorm.DB
}

func NewGoodsDao() *GoodsDao {
	return &GoodsDao{db: mygorm.DefaultDB}
}

func (this *GoodsDao) InsertGoods(goods model.Goods) error {
	return this.db.Create(&goods).Error
}
