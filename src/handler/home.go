package handler

import (
	"github.com/gasolzhang/jdbackend/src/consts"
	"github.com/gasolzhang/jdbackend/src/info"
	"github.com/gasolzhang/jdbackend/src/response"
	"github.com/gasolzhang/jdbackend/src/service"
	"github.com/gin-gonic/gin"
)

func GetHomeBanner(ctx *gin.Context) {
	var form struct {
		Limit int `validate:"required"`
	}
	var err error

	context := response.MyContext{Context: ctx}
	err = context.ShouldBind(&form)
	if err != nil {
		context.ResponseParamError()
		return
	}

	homeService := service.NewHomeService()
	banner, err := homeService.GetBanner(form.Limit)
	if err != nil {
		context.ResponseCustomError(consts.ErrorCodeServer, err.Error())
		return
	}

	var resp struct {
		Banner *[]info.HomeBanner
	}
	resp.Banner = &banner

	context.ResponseData(resp)
}

func InsertHomeBanner(ctx *gin.Context) {
	var form struct {
		Image  string `validate:"required,url"`
		Action string `validate:"omitempty"`
	}
	var err error

	context := response.MyContext{Context: ctx}
	err = context.ShouldBind(&form)
	if err != nil {
		context.ResponseParamError()
		return
	}

	homeService := service.NewHomeService()
	err = homeService.InsertBanner(form.Image, form.Action)
	if err != nil {
		context.ResponseCustomError(consts.ErrorCodeServer, err.Error())
		return
	}
	context.ResponseOK()
}

func GetHomeGuess(ctx *gin.Context) {
	var form struct {
		Limit int `validate:"required"`
	}

	var err error

	context := response.MyContext{Context: ctx}
	err = context.ShouldBind(&form)
	if err != nil {
		context.ResponseParamError()
		return
	}

	homeService := service.NewHomeService()
	guess, err := homeService.GetGuess(form.Limit)
	if err != nil {
		context.ResponseCustomError(consts.ErrorCodeServer, err.Error())
		return
	}
	var resp struct{
		Guess *[]info.HomeGuess
	}
	resp.Guess = &guess
	context.ResponseData(resp)
}

func InsertHomeGuess(ctx *gin.Context) {
	var form struct {
		Image   string `validate:"required,url"`
		GoodsId string `validate:"required"`
		Price   string `validate:"required"`
	}
	var err error

	context := response.MyContext{Context: ctx}
	err = context.ShouldBind(&form)
	if err != nil {
		context.ResponseParamError()
		return
	}

	homeService := service.NewHomeService()
	err = homeService.InsertGuess(form.Image, form.GoodsId, form.Price)
	if err != nil {
		context.ResponseCustomError(consts.ErrorCodeServer, err.Error())
		return
	}
	context.ResponseOK()
}

func GetHotGoods(ctx *gin.Context) {
	var form struct {
		Page  int `validate:"required,min=1"`
		Limit int `validate:"omitempty,omitempty=20"`
	}

	var err error

	context := response.MyContext{Context: ctx}
	err = context.ShouldBind(&form)
	if err != nil {
		context.ResponseParamError()
		return
	}

	homeService := service.NewHomeService()
	goods, err := homeService.GetHotGoods(form.Page, form.Limit)
	if err != nil {
		context.ResponseCustomError(consts.ErrorCodeServer, err.Error())
		return
	}

	var resp struct{
		Goods *[]info.Goods
	}
	resp.Goods = &goods

	context.ResponseData(resp)
}
