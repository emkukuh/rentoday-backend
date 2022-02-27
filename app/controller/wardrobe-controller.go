package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mashingan/smapping"

	"rentoday.id/app/constant"
	"rentoday.id/app/dto"
	"rentoday.id/app/helper"
	"rentoday.id/app/service"
)

var (
	wardrobeService = service.NewWardrobeService()
)

type WardrobeController interface {
	FindAll(ctx *gin.Context)
	FindByUser(ctx *gin.Context)
	AddWardrobe(ctx *gin.Context)
}

type wardrobeController struct{}

func NewWardrobeController() WardrobeController {
	return &wardrobeController{}
}

func (c *wardrobeController) FindAll(ctx *gin.Context) {
	wardrobes, err := wardrobeService.FindAll()
	if err != nil {
		response := helper.BuildErrorResponse(constant.ErrorRequestMessage, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	response := helper.BuildSuccessResponse(wardrobes)
	ctx.JSON(http.StatusOK, response)
}

func (c *wardrobeController) FindByUser(ctx *gin.Context) {
	userId, err := getUserIdInHeaderToken(ctx)
	if err != nil {
		response := helper.BuildErrorResponse("user not found", err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}
	wardrobes, err := wardrobeService.GetListByUserId(userId)
	if err != nil {
		response := helper.BuildErrorResponse(constant.ErrorRequestMessage, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	response := helper.BuildSuccessResponse(wardrobes)
	ctx.JSON(http.StatusOK, response)
}

func (c *wardrobeController) AddWardrobe(ctx *gin.Context) {
	var wardrobeDto dto.AddWardrobeRequest
	ctx.BindJSON(&wardrobeDto)
	authHeader := helper.GetHeaderAuth(ctx)
	userId, err := jwtService.GetUserIdByToken(authHeader)
	if err == nil {
		wardrobeDto.UserID = userId
	}
	res, err := wardrobeService.Create(wardrobeDto)
	if err != nil {
		response := helper.BuildErrorResponse(constant.ErrorRequestMessage, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	var wardrobeResponse dto.AddWardrobeResponse
	smapping.FillStruct(&wardrobeResponse, smapping.MapFields(&res))
	response := helper.BuildSuccessResponse(wardrobeResponse)
	ctx.JSON(http.StatusOK, response)
}
