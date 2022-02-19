package auth

import (
	"github.com/gin-gonic/gin"
	"rentoday.id/app/controller"
)

var (
	authController 	controller.AuthControllerInterface = controller.NewAuthController()
)

func CreateRouter(router *gin.Engine) {
	group := router.Group("api/auth")
	group.POST("/login", authController.LoginUser)
	group.POST("/register", authController.RegisterUser)
	group.POST("/admin/register", authController.RegisterAdmin)
	group.POST("/admin/login", authController.LoginAdmin)
}