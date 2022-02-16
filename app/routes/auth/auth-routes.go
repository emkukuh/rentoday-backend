package auth

import (
	"github.com/gin-gonic/gin"
	"rentoday.id/app/controller"
)

var (
	authController 	controller.AuthControllerInterface = controller.NewAuthController()
)

func postLogin(ctx *gin.Context) {
	authController.Login(ctx)
}

func postRegister(ctx *gin.Context) {
	authController.Register(ctx)
}

func CreateRouter(router *gin.Engine) {
	group := router.Group("api/auth")
	group.POST("/login", postLogin)
	group.POST("/register", postRegister)
}