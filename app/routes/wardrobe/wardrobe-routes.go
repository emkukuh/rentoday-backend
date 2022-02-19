package wardrobe

import (
	"github.com/gin-gonic/gin"
	"rentoday.id/app/controller"
	"rentoday.id/app/middleware"
)

var (
	wardrobeController controller.WardrobeController = controller.NewWardrobeController()
)

func CreateRouter(router *gin.Engine) {
	group := router.Group("api/wardrobe", middleware.AuthJwt())
	group.GET("/list", wardrobeController.FindAll)
	group.POST("/add", wardrobeController.AddWardrobe)
}
