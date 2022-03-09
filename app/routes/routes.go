package routes

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"rentoday.id/app/routes/auth"
	"rentoday.id/app/routes/parameter"
	"rentoday.id/app/routes/wardrobe"
)

func SetupRoutes(router *gin.Engine) *gin.Engine {
	setupCors(router)
	auth.CreateRouter(router)
	parameter.CreateRouter(router)
	wardrobe.CreateRouter(router)
	return router
}

func setupCors(router *gin.Engine) {
	router.Use(cors.New(cors.Config{
		AllowOrigins:  []string{"http://localhost:3000"},
		AllowMethods:  []string{"GET", "POST", "PATCH", "OPTIONS"},
		AllowHeaders:  []string{"Authorization"},
		ExposeHeaders: []string{"Content-Length"},
	}))
}
