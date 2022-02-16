package configs

import (
	"github.com/gin-gonic/gin"
	"rentoday.id/app/controller"
)

var configController = controller.NewConfigController()

func CreateRouter(router *gin.Engine) {
	group := router.Group("api/config")
	group.GET("/wardrobe-categories", configController.GetWardrobeCategories)
}