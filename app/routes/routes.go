package routes

import (
	"github.com/gin-gonic/gin"
	"rentoday.id/app/routes/auth"
	"rentoday.id/app/routes/configs"
	"rentoday.id/app/routes/wardrobe"
)

func SetupRoutes(router *gin.Engine) *gin.Engine {
	wardrobe.CreateRouter(router)
	auth.CreateRouter(router)
	configs.CreateRouter(router)
	return router
}