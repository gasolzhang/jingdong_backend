package main

import (
	"github.com/gasolzhang/jdbackend/src/mygin"
	"github.com/gasolzhang/jdbackend/src/mygorm"
	"log"
)

func main() {
	db := mygorm.Init()
	if db != nil {
		defer db.Close()
		mygin.Init()
	} else {
		log.Println("数据库初始化失败")
	}

}
