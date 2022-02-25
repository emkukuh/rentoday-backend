package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"rentoday.id/app/helper"
	"rentoday.id/app/service"
)

var parameterService = service.NewParameterService()

type ParameterController interface {
	GetWardrobeCategories(ctx *gin.Context)
	GetWardrobeMaterials(ctx *gin.Context)
}

type parameterController struct{}

func NewParameterController() ParameterController {
	return &parameterController{}
}

func (c *parameterController) GetWardrobeCategories(ctx *gin.Context) {
	categories, err := parameterService.GetWardrobeCategories()
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, nil)
		return
	}
	response := helper.BuildResponse(true, categories)
	ctx.JSON(http.StatusOK, response)
}

func (c *parameterController) GetWardrobeMaterials(ctx *gin.Context) {
	materials, err := parameterService.GetWardrobeMaterials()
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, nil)
		return
	}
	response := helper.BuildResponse(true, materials)
	ctx.JSON(http.StatusOK, response)
}
