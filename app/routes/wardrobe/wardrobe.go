package wardrobe

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"rentoday.id/app/controller"
	"rentoday.id/app/service"
)

var (
	wardrobeServices service.WardrobeServiceInterface	= service.New()
	wardrobeController controller.WardrobeController	= controller.New(wardrobeServices)
)

func postWardrobe(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, wardrobeController.Save(ctx))
}

func getWardrobeList(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, wardrobeController.FindAll())
}

func CreateRouter(router *gin.Engine) {
	group := router.Group("wardrobe")
	group.GET("/list", getWardrobeList)	
	group.POST("/add", postWardrobe)	
}