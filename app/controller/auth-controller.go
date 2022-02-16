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
		response := helper.BuildErrorResponse("failed to process request", err.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	authResult := authService.VerifyCredential(loginDto.Email, loginDto.Password)
	if v, ok := authResult.(model.User); ok {
		generatedToken := jwtService.GenerateToken(strconv.FormatUint(v.ID, 10))
		v.Token = generatedToken
		response := helper.BuildResponse(true, v)
		ctx.JSON(http.StatusOK,response)
		return
	}
	response := helper.BuildErrorResponse("invalid credential", "invalid credential", helper.EmptyObj{})
	ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
}

func (c *authController) Register(ctx *gin.Context) {
	var registerDto dto.RegisterDto
	err := ctx.ShouldBind(&registerDto)
	if err != nil {
		response := helper.BuildErrorResponse(constant.ErrorRequestMessage, err.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	newUser, newUserErr := authService.Register(registerDto)
	if newUserErr != nil {
		response := helper.BuildErrorResponse(constant.ErrorRequestMessage, newUserErr.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	response := helper.BuildResponse(true, newUser)

	// isDuplicated, err := authService.IsDuplicateEmail(registerDto.Email)
	// if err != nil {
	// 	response := helper.BuildErrorResponse(constant.ErrorRequestMessage, err.Error(), helper.EmptyObj{})
	// 	ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
	// 	return
	// }
	// if isDuplicated {
	// 	response := helper.BuildErrorResponse(constant.ErrorRequestMessage, "email duplicate", helper.EmptyObj{})
	// 	ctx.JSON(http.StatusConflict, response)
	// 	return
	// }
	// authService.CreateUser(registerDto)
	ctx.JSON(http.StatusOK, response)

	// if isDuplicted, err := c.authService.IsDuplicateEmail(registerDto.Email) {
	// 	response := helper.BuildErrorResponse(constant.ErrorRequestMessage, "email duplicate", helper.EmptyObj{})
	// 	ctx.JSON(http.StatusConflict, response)
	// 	return
	// } else {
		// createdUser := c.authService.CreateUser(registerDto)
		// token := c.jwtService.GenerateToken(strconv.FormatUint(createdUser.ID, 10))
		// createdUser.Token = token
		// response := helper.BuildResponse(true, createdUser)
		// ctx.JSON(http.StatusCreated, response)
	// 	ctx.JSON(http.StatusOK, nil)
	// 	return
	// }
}