package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"rentoday.id/app/helper"
	"rentoday.id/app/service"
)

var configService = service.NewConfigService()

type ConfigController interface {
	GetWardrobeCategories(ctx *gin.Context)
}

type configController struct {}

func NewConfigController() ConfigController {
	return &configController{}
}

func (c *configController) GetWardrobeCategories(ctx *gin.Context) {
	categories, err := configService.GetWardrobeCategories()
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, nil)
		return
	}
	response := helper.BuildResponse(true, categories)
	ctx.JSON(http.StatusOK, response)
}