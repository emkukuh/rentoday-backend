package routes

import (
	"github.com/gin-gonic/gin"
	"rentoday.id/app/routes/auth"
	"rentoday.id/app/routes/parameter"
	"rentoday.id/app/routes/wardrobe"
)

func SetupRoutes(router *gin.Engine) *gin.Engine {
	auth.CreateRouter(router)
	parameter.CreateRouter(router)
	wardrobe.CreateRouter(router)
	return router
}
