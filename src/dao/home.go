package dao

import (
	"github.com/gasolzhang/jdbackend/src/info"
	"github.com/gasolzhang/jdbackend/src/model"
	"github.com/gasolzhang/jdbackend/src/mygorm"
	"github.com/jinzhu/gorm"
)

type HomeDao struct {
	db *gorm.DB
}

func NewHomeDao() *HomeDao {
	return &HomeDao{db: mygorm.DefaultDB}
}

func (this *HomeDao) GetBanner(limit int) (banner []info.HomeBanner, err error) {
	db := this.db.Table("home_banner").Select("image,action").Limit(limit).Scan(&banner)
	return banner, db.Error
}

func (this *HomeDao) InsertBanner(banner *model.HomeBanner) error {
	return this.db.Create(banner).Error
}

func (this *HomeDao) InsertGuess(guess *model.HomeGuess) error {
	return this.db.Create(guess).Error
}

func (this *HomeDao) GetGuess(limit int) (guess []info.HomeGuess, err error) {
	db := this.db.Table("home_guess").Select("image,goods_id,price").Limit(limit).Scan(&guess)
	return guess, db.Error
}

func (this *HomeDao) GetHotGoods(page, limit int) (goods []info.Goods, err error) {
	db := this.db.Table("goods").
		Select("id,cover,name,price,old_price,cate_id,sale_count").
		Where("sale_count > ?", 100).
		Offset((page - 1) * limit).
		Limit(limit).
		Order("sale_count DESC").Scan(&goods)
	return goods, db.Error
}


