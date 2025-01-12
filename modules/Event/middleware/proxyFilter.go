package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ProxyFilterAll(ipWL string) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {

		if ctx.RemoteIP() != ipWL {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized", "message": "Client not define", "status": false, "data": nil})
			return
		}
	}
}
