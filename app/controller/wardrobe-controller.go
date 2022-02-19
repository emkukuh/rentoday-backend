package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"rentoday.id/app/constant"
	"rentoday.id/app/dto"
	"rentoday.id/app/helper"
	"rentoday.id/app/model"
	"rentoday.id/app/service"
)

var wardrobeService = service.NewWardrobeService()

type WardrobeController interface {
	FindAll() []model.Wardrobe
	AddWardrobe(ctx *gin.Context)
}

type wardrobeController struct{}

func NewWardrobeController() WardrobeController {
	return &wardrobeController{}
}

func (c *wardrobeController) FindAll() []model.Wardrobe {
	return wardrobeService.FindAll()
}

func (c *wardrobeController) AddWardrobe(ctx *gin.Context) {
	var wardrobeDto dto.AddWardrobe
	ctx.BindJSON(&wardrobeDto)
	res, err := wardrobeService.Create(wardrobeDto)
	if err != nil {
		response := response.BuildErrorResponse(constant.ErrorRequestMessage, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	ctx.JSON(http.StatusOK, res)
}
