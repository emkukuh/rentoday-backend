package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"rentoday.id/app/constant"
	"rentoday.id/app/dto"
	"rentoday.id/app/helper"
	"rentoday.id/app/model"
	"rentoday.id/app/service"
)

var jwtService service.JwtServiceInterface = service.NewJwtService()
var authService = service.NewAuthService()

type AuthControllerInterface interface {
	Login(ctx *gin.Context)
	Register(ctx *gin.Context)
}

type authController struct {}

func NewAuthController() AuthControllerInterface {
	return &authController{}
}

func (c *authController) Login(ctx *gin.Context) {
	var loginDto dto.LoginDto
	err := ctx.ShouldBind(&loginDto)
	if err != nil {
		response := response.BuildErrorResponse("failed to process request", err.Error(), response.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	authResult := authService.VerifyCredential(loginDto.Email, loginDto.Password)
	if v, ok := authResult.(model.User); ok {
		generatedToken := jwtService.GenerateToken(strconv.FormatUint(v.ID, 10))
		v.Token = generatedToken
		response := response.BuildResponse(true, v)
		ctx.JSON(http.StatusOK,response)
		return
	}
	response := response.BuildErrorResponse("invalid credential", "invalid credential", response.EmptyObj{})
	ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
}

func (c *authController) Register(ctx *gin.Context) {
	var registerDto dto.RegisterDto
	err := ctx.ShouldBind(&registerDto)
	if err != nil {
		response := response.BuildErrorResponse(constant.ErrorRequestMessage, err.Error(), response.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	newUser, newUserErr := authService.Register(registerDto)
	if newUserErr != nil {
		response := response.BuildErrorResponse(constant.ErrorRequestMessage, newUserErr.Error(), response.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	response := response.BuildResponse(true, newUser)

	ctx.JSON(http.StatusOK, response)
}