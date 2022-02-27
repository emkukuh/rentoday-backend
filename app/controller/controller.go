package controller

import (
	"github.com/gin-gonic/gin"
	"rentoday.id/app/helper"
	"rentoday.id/app/service"
)

var jwtService service.JwtServiceInterface = service.NewJwtService()

func getUserIdInHeaderToken(ctx *gin.Context) (string, error) {
	authHeader := helper.GetHeaderAuth(ctx)
	aToken, err := jwtService.GetUserIdByToken(authHeader)
	return aToken, err
}
