package app

import (
	"rentoday.id/app/database"
	"rentoday.id/app/routes"

	"github.com/gin-gonic/gin"
)

func Run() {
	router := gin.Default()
	database.Start()
	routes.SetupRoutes(router)
	router.Run(":4000")
}