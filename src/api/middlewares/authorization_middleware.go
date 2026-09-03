package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CheckBeareToken() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		bearerToken := ctx.GetHeader("Authorization")
		if bearerToken == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"result": "No Authorization has set.",
			})
		}
		ctx.Next()
	}
}
