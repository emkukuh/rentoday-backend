package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mashingan/smapping"
	"rentoday.id/app/constant"
	"rentoday.id/app/dto"
	"rentoday.id/app/helper"
	"rentoday.id/app/model"
	"rentoday.id/app/service"
)

var (
	jwtService       service.JwtServiceInterface = service.NewJwtService()
	authUserService                              = service.AuthUserService
	authAdminService                             = service.AuthAdminService
)

type AuthControllerInterface interface {
	LoginUser(ctx *gin.Context)
	LoginAdmin(ctx *gin.Context)
	RegisterUser(ctx *gin.Context)
	RegisterAdmin(ctx *gin.Context)
}

type authController struct{}

func NewAuthController() AuthControllerInterface {
	return &authController{}
}

func (c *authController) LoginUser(ctx *gin.Context) {
	var loginDto dto.LoginDto
	err := ctx.ShouldBind(&loginDto)
	if err != nil {
		handleBadRequestError(ctx, err)
		return
	}
	authResult := authUserService.VerifyCredential(loginDto.Email, loginDto.Password)
	if v, ok := authResult.(model.User); ok {
		generatedToken := jwtService.GenerateToken(strconv.FormatUint(v.ID, 10))
		v.AccessToken = generatedToken
		response := response.BuildResponse(true, v)
		ctx.JSON(http.StatusOK, response)
		return
	}
	handleInvalidTokenError(ctx)
}

func (c *authController) LoginAdmin(ctx *gin.Context) {
	var loginDto dto.LoginDto
	err := ctx.ShouldBind(&loginDto)
	if err != nil {
		handleBadRequestError(ctx, err)
		return
	}
	authResult, err := authAdminService.VerifyCredential(loginDto.Email, loginDto.Password)
	if err != nil {
		handleInvalidTokenError(ctx)
		return
	}
	generatedToken := jwtService.GenerateToken(strconv.FormatUint(authResult.ID, 10))
	authResult.AccessToken = generatedToken
	var loginResponseDto dto.LoginAdminResponseDto
	smapping.FillStruct(&loginResponseDto, smapping.MapFields(&authResult))
	loginResponseDto.ID = authResult.ID
	response := response.BuildSuccessResponse(loginResponseDto)
	ctx.JSON(http.StatusOK, response)
}

func (c *authController) RegisterUser(ctx *gin.Context) {
	var registerDto dto.RegisterDto
	err := ctx.ShouldBind(&registerDto)
	if err != nil {
		handleBadRequestError(ctx, err)
		return
	}
	newUser, err := authUserService.Register(registerDto)
	if err != nil {
		handleBadRequestError(ctx, err)
		return
	}
	response := response.BuildResponse(true, newUser)
	ctx.JSON(http.StatusOK, response)
}

func (c *authController) RegisterAdmin(ctx *gin.Context) {
	var registerDto dto.RegisterDto
	err := ctx.ShouldBind(&registerDto)
	if err != nil {
		handleBadRequestError(ctx, err)
		return
	}
	newAdmin, err := authAdminService.Register(registerDto)
	if err != nil {
		handleBadRequestError(ctx, err)
		return
	}
	response := response.BuildResponse(true, newAdmin)
	ctx.JSON(http.StatusOK, response)
}

func handleBadRequestError(ctx *gin.Context, err error) {
	response := response.BuildErrorResponse(constant.ErrorRequestMessage, err.Error(), response.EmptyObj{})
	ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
}

func handleInvalidTokenError(ctx *gin.Context) {
	response := response.BuildErrorResponse("invalid credential", "invalid credential", response.EmptyObj{})
	ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
}
