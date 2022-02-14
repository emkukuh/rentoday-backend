package app

import (
	"rentoday.id/app/controller"
	"rentoday.id/app/database"
	"rentoday.id/app/service"

	"github.com/gin-gonic/gin"
)

var (
	wardrobeServices service.WardrobeServiceInterface	= service.New()
	wardrobeController controller.WardrobeController	= controller.New(wardrobeServices)
)

func Run() {
	router := gin.Default()
	database.Start()
	router.GET("/wardrobe/list", getWardrobeList)
	router.POST("/wardrobe/add", postWardrobe)
	router.Run(":4000")
}

func getWardrobeList(ctx *gin.Context) {
	ctx.JSON(200, wardrobeController.FindAll())
}

func postWardrobe(ctx *gin.Context) {
	ctx.JSON(200, wardrobeController.Save(ctx))
}