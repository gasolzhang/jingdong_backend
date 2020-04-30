package middleware

import (
	"github.com/gasolzhang/jdbackend/src/consts"
	"github.com/gasolzhang/jdbackend/src/jwt"
	"github.com/gin-gonic/gin"
	"log"
)

func CheckToken() gin.HandlerFunc {
	return func(context *gin.Context) {
		token := context.GetHeader("Token")
		_, err := jwt.Decode(token)
		if err != nil {
			log.Printf("token解析失败：%v", err)
			context.AbortWithStatusJSON(consts.ErrorCodeTokenError, gin.H{"message": "token错误或已失效"})
		}
	}
}
