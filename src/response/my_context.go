package response

import (
	"github.com/gasolzhang/jdbackend/src/consts"
	"github.com/gin-gonic/gin"
	"net/http"
)

type messageResponse struct {
	Message string
	Code    int
}

type dataResponse struct {
	messageResponse
	Data interface{}
}

type MyContext struct {
	*gin.Context
}

func (this *MyContext) ResponseOK() {
	response := new(messageResponse)
	response.Message = "请求成功"
	response.Code = consts.ErrorCodeOk
	this.Context.JSON(http.StatusOK, response)
}

func (this *MyContext) ResponseParamError() {
	response := new(messageResponse)
	response.Message = "请求参数错误"
	response.Code = consts.ErrorCodeParameter
	this.Context.JSON(http.StatusOK, response)
}

func (this *MyContext) ResponseCustomError(code int, message string) {
	response := new(messageResponse)
	response.Message = message
	response.Code = code
	this.JSON(http.StatusOK, response)
}

func (this *MyContext) ResponseData(obj interface{}) {
	response := new(dataResponse)
	response.Message = "请求成功"
	response.Data = obj
	response.Code = consts.ErrorCodeOk
	this.Context.JSON(http.StatusOK, response)
}
