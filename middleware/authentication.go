package middleware

import (
	"golang-pinjol/helper"
	"golang-pinjol/services"
	"log"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func Authorize(jwtService services.JwtService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		header := ctx.GetHeader("Authorization")
		if header == "" {
			reponse := helper.ErrorResponse("failed to Authorization ", "token not found", nil)
			ctx.AbortWithStatusJSON(http.StatusBadRequest, reponse)
			return
		}
		token, err := jwtService.ValidateTokenService(header)
		if err != nil {
			log.Println(err)
			response := helper.ErrorResponse("your token is invalid ", err.Error(), nil)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		if token.Valid {
			claims := token.Claims.(jwt.MapClaims)
			log.Println("claims[customer_id] :", claims["customer_id"])
			log.Println("claims[issuer] :", claims["issuer"])
		} else {
			response := helper.ErrorResponse("your token is invalid ", "invalid token", nil)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
	}
}
