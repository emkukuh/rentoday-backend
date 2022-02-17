package parameter

import (
	"github.com/gin-gonic/gin"
	"rentoday.id/app/controller"
)

var parameterController = controller.NewParameterController()

func CreateRouter(router *gin.Engine) {
	group := router.Group("api/parameter")
	group.GET("/wardrobe-categories", parameterController.GetWardrobeCategories)
	group.GET("/wardrobe-materials", parameterController.GetWardrobeMaterials)
}