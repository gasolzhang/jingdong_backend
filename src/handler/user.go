package handler

import (
	"github.com/gasolzhang/jdbackend/src/consts"
	"github.com/gasolzhang/jdbackend/src/response"
	"github.com/gasolzhang/jdbackend/src/service"
	"github.com/gin-gonic/gin"
)

func Login(ctx *gin.Context) {
	var form struct {
		Phone    string `validate:"required,len=11"`
		Password string `validate:"required,min=6,max=12"`
	}

	var err error

	context := response.MyContext{Context: ctx}
	err = ctx.ShouldBind(&form)
	if err != nil {
		context.ResponseParamError()
		return
	}

	userService := service.NewUserService()
	err, code := userService.Login(form.Phone, form.Password)
	if err != nil {
		context.ResponseCustomError(code, err.Error())
		return
	}

	user, err := userService.GetUserInfo(form.Phone)
	if err != nil {
		context.ResponseCustomError(consts.ErrorCodeServer, err.Error())
		return
	}
	//发token
	user.Token = userService.GenToken(form.Phone)
	context.ResponseData(user)
}

func Regist(ctx *gin.Context) {
	var form struct {
		Phone    string `validate:"required,len=11"`
		Password string `validate:"required,min=6,max=12"`
		Code     string `validate:"required,len=6"`
	}

	var err error

	context := response.MyContext{Context: ctx}
	err = context.ShouldBind(&form)
	if err != nil {
		context.ResponseParamError()
		return
	}

	userService := service.NewUserService()
	err, errCode := userService.Regist(form.Phone, form.Password, form.Code)
	if err != nil {
		context.ResponseCustomError(errCode, err.Error())
		return
	}

	user, err := userService.GetUserInfo(form.Phone)
	if err != nil {
		context.ResponseCustomError(consts.ErrorCodeServer, err.Error())
		return
	}
	//发token
	user.Token = userService.GenToken(form.Phone)
	context.ResponseData(user)
}

func ResetPassword(ctx *gin.Context) {
	var form struct {
		Phone    string `validate:"required,len=11"`
		Password string `validate:"required,min=6,max=12"`
		Code     string `validate:"required,len=6"`
	}

	var err error

	context := response.MyContext{Context: ctx}
	err = context.ShouldBind(&form)
	if err != nil {
		context.ResponseParamError()
		return
	}

	userService := service.NewUserService()
	err, errCode := userService.ResetPassword(form.Phone, form.Password, form.Code)
	if err != nil {
		context.ResponseCustomError(errCode, err.Error())
		return
	}
	context.ResponseOK()
}

func ModifyPassword(ctx *gin.Context) {
	var form struct {
		Phone    string `validate:"required,len=11"`
		Password string `validate:"required,min=6,max=12"`
	}

	var err error

	context := response.MyContext{Context: ctx}
	err = context.ShouldBind(&form)
	if err != nil {
		context.ResponseParamError()
		return
	}

	userService := service.NewUserService()
	err, errCode := userService.ModifyPassword(form.Phone, form.Password)
	if err != nil {
		context.ResponseCustomError(errCode, err.Error())
		return
	}
	context.ResponseOK()
}
