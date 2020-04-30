package mygorm

import _ "github.com/jinzhu/gorm/dialects/mysql"
import (
	"fmt"
	"github.com/gasolzhang/jdbackend/src/model"
	"github.com/jinzhu/gorm"
)

var DefaultDB *gorm.DB

func Init() *gorm.DB {
	//在全部表前面加前缀
	//gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
	//	return "custom"+defaultTableName
	//}
	db, err := gorm.Open("mysql", "root:123456@tcp(localhost:3306)/jdbackend?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Printf("数据库连接失败，err=%v", err)
		return nil
	}
	// 关闭复数表名，如果设置为true，`User`表的表名就会是`user`，而不是`users`
	db.SingularTable(true)
	db.LogMode(true)

	migrateTables(db)

	DefaultDB = db

	return db
}

func migrateTables(db *gorm.DB) {
	//如果有则更新，没有则创建
	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.HomeBanner{})
	db.AutoMigrate(&model.HomeGuess{})
	db.AutoMigrate(&model.Goods{})
	//exists := db.HasTable(&model.User{})
	//if !exists {
	//	db.CreateTable(&model.User{})
	//	//db.Table("自定义tablename").CreateTable(&model.User{})
	//	fmt.Println("创建了数据表")
	//} else {
	//	db.AutoMigrate(&model.User{})
	//}
}
