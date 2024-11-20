package middleware

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// ApiKeyAuthMiddleware authenticates requests based on API keys
func ApiKeyAuthMiddleware(allowedKeys map[string]bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		apiKey := c.GetHeader("API-KEY")
		if _, exists := allowedKeys[apiKey]; !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
			log.Panicf("Unauthorized Access")
			c.Abort()
			return
		}
		c.Next()
	}
}
