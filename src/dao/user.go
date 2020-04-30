package dao

import (
	"github.com/gasolzhang/jdbackend/src/info"
	"github.com/gasolzhang/jdbackend/src/model"
	"github.com/gasolzhang/jdbackend/src/mygorm"
	"github.com/jinzhu/gorm"
)

type UserDao struct {
	db *gorm.DB
}

func NewUserDao() *UserDao {
	dao := UserDao{db: mygorm.DefaultDB}
	return &dao
}

func (this *UserDao) Login(phone, password string) error {
	return this.db.Where("phone = ? AND password = ?", phone, password).First(new(model.User)).Error
}

func (this *UserDao) Regist(phone, password string) error {
	user := new(model.User)
	user.Name = "用户" + phone
	user.Phone = phone
	user.Password = password
	return this.db.Create(user).Error
}

func (this *UserDao) ResetPassword(phone, password string) error {
	return this.db.Model(new(model.User)).Where("phone = ?", phone).Updates(model.User{Password: password}).Error
}

func (this *UserDao) GetUserInfo(phone string) (user info.User, err error) {
	db := this.db.Table("user").Select("id,phone,password,name,avatar,sign").Where("phone = ?", phone).Scan(&user)
	return user, db.Error
}
