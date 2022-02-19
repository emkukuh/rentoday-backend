package wardrobe

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"rentoday.id/app/controller"
	"rentoday.id/app/middleware"
)

var (
	wardrobeController controller.WardrobeController = controller.NewWardrobeController()
)

func getWardrobeList(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, wardrobeController.FindAll())
}

func CreateRouter(router *gin.Engine) {
	group := router.Group("api/wardrobe", middleware.AuthJwt())
	group.GET("/list", getWardrobeList)
	group.POST("/add", wardrobeController.AddWardrobe)
}
