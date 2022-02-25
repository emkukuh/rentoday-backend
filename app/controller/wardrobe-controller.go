package controller

import (
	"log"
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
	// var wardrobesResponse []dto.AddWardrobeResponse
	// for _, wardrobe := range wardrobes {
	// 	var wardrobeRes dto.AddWardrobeResponse
	// 	smapping.FillStruct(&wardrobeRes, smapping.MapFields(&wardrobe))
	// 	wardrobesResponse = append(wardrobesResponse, wardrobeRes)
	// }
	// if err != nil {
	// 	response := helper.BuildErrorResponse("error maaping", err.Error(), nil)
	// 	ctx.AbortWithStatusJSON(http.StatusInternalServerError, response)
	// 	return
	// }
	response := helper.BuildSuccessResponse(wardrobes)
	ctx.JSON(http.StatusOK, response)
}

func (c *wardrobeController) AddWardrobe(ctx *gin.Context) {
	var wardrobeDto dto.AddWardrobeRequest
	ctx.BindJSON(&wardrobeDto)
	log.Print("==============")
	log.Print(wardrobeDto)
	authHeader := helper.GetHeaderAuth(ctx)
	userId, err := jwtService.GetUserIdByToken(authHeader)
	if err == nil {
		wardrobeDto.UserID = userId
	}
	log.Print("user id ====== ", userId)
	res, err := wardrobeService.Create(wardrobeDto)
	if err != nil {
		response := helper.BuildErrorResponse(constant.ErrorRequestMessage, err.Error(), nil)
		log.Print("====== ERROR")
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	var wardrobeResponse dto.AddWardrobeResponse
	smapping.FillStruct(&wardrobeResponse, smapping.MapFields(&res))
	response := helper.BuildSuccessResponse(wardrobeResponse)
	ctx.JSON(http.StatusOK, response)
}
