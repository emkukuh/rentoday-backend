package wardrobe

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"rentoday.id/app/controller"
)

var (
	wardrobeController controller.WardrobeController	= controller.NewWardrobeController()
)

func postWardrobe(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, wardrobeController.Save(ctx))
}

func getWardrobeList(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, wardrobeController.FindAll())
}

func CreateRouter(router *gin.Engine) {
	group := router.Group("api/wardrobe")
	group.GET("/list", getWardrobeList)	
	group.POST("/add", postWardrobe)	
}