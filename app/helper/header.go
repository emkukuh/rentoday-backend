package helper

import (
	"github.com/gin-gonic/gin"
	"rentoday.id/app/constant"
)

func GetHeaderAuth(ctx *gin.Context) string {
	auth := ctx.GetHeader(constant.Auth)
	return auth
}
