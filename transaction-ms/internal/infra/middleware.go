package infra

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func JSONBodyMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if ctx.GetHeader("Content-Type") != "application/json" {
			ctx.Abort()
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "this route expects Content-Type with application/json"})
		}
	}
}
