package routes

import (
	"github.com/gin-gonic/gin"
	"rentoday.id/routes/wardrobe"
)

func SetupRoutes(router *gin.Engine) *gin.Engine {
	wardrobe.CreateRouter(router)
	return router
}