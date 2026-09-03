package middlewares

import (
	"net/http"

	"github.com/didip/tollbooth"
	"github.com/gin-gonic/gin"
)

func RateLimitter() gin.HandlerFunc {
	lmt := tollbooth.NewLimiter(1, nil)
	return func(ctx *gin.Context) {
		err := tollbooth.LimitByRequest(lmt, ctx.Writer, ctx.Request)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{
				"result": err.Error(),
			})
		} else {
			ctx.Next()
		}
	}
}
