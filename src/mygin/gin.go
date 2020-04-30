package mygin

import (
	"github.com/gasolzhang/jdbackend/src/handler"
	"github.com/gasolzhang/jdbackend/src/middleware"
	"github.com/gin-gonic/gin"
)

func Init() {
	engine := gin.Default()
	//engine.Use()
	rootGroup := engine.Group("/v1/api")

	userGroup := rootGroup.Group("/user")
	{
		userGroup.POST("/login", handler.Login)
		userGroup.POST("/regist", handler.Regist)
		userGroup.POST("/reset_password", handler.ResetPassword)
		userGroup.POST("/modify_password", middleware.CheckToken(), handler.ModifyPassword)
	}

	homeGroup := rootGroup.Group("/home")
	{
		homeGroup.GET("/banner", handler.GetHomeBanner)
		homeGroup.POST("/insert_banner", handler.InsertHomeBanner)
		homeGroup.GET("/guess", handler.GetHomeGuess)
		homeGroup.POST("/insert_guess", handler.InsertHomeGuess)
		homeGroup.GET("/hot_goods", handler.GetHotGoods)
	}

	goodsGroup := rootGroup.Group("/goods")
	{
		goodsGroup.POST("/insert", handler.InsertGoods)
	}

	engine.Run(":8080")
}
