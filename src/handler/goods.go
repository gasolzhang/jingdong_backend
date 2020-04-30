package handler

import (
	"github.com/gasolzhang/jdbackend/src/consts"
	"github.com/gasolzhang/jdbackend/src/model"
	"github.com/gasolzhang/jdbackend/src/response"
	"github.com/gasolzhang/jdbackend/src/service"
	"github.com/gin-gonic/gin"
)

func InsertGoods(ctx *gin.Context) {
	var form struct {
		Cover     string `validate:"required,url"`
		Name      string `validate:"required"`
		Price     string `validate:"required"`
		OldPrice  string `validate:"omitempty"`
		CateId    string `validate:"required"`
		SaleCount int    `validate:"omitempty"`
	}

	var err error

	context := response.MyContext{Context: ctx}
	err = context.ShouldBind(&form)
	if err != nil {
		context.ResponseParamError()
		return
	}

	goodsService := service.NewGoodsService()
	goods := model.Goods{Cover: form.Cover, Name: form.Name, Price: form.Price, OldPrice: form.OldPrice, CateId: form.CateId, SaleCount: form.SaleCount}
	err = goodsService.InsertGoods(goods)
	if err != nil {
		context.ResponseCustomError(consts.ErrorCodeServer, err.Error())
		return
	}

	context.ResponseOK()
}
