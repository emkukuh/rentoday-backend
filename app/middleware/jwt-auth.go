package middleware

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"rentoday.id/app/constant"
	"rentoday.id/app/helper"
	"rentoday.id/app/service"
)

func AuthJwt(jwtService service.JwtServiceInterface) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader(constant.Auth)
		if authHeader == "" {
			response := helper.BuildErrorResponse("Failed to process request", "Token not found", nil)
			c.AbortWithStatusJSON(http.StatusBadRequest, response)
		}
		token, err := jwtService.ValidateToken(authHeader)
		if token.Valid {
			claims := token.Claims.(jwt.MapClaims)
			log.Println("Claim[email]: ", claims["email"])
			log.Println("Claim[issuer]: ", claims["issuer"])
		} else {
			log.Println(err)
			response := helper.BuildErrorResponse("Token not valid", err.Error(), nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
		}
	}
}