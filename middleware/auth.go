package middleware

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Map of allowed API keys
var apiKeys = map[string]bool{
	"cs218secret": true,
}

// ApiKeyAuthMiddleware authenticates requests based on API keys
func ApiKeyAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		apiKey := c.GetHeader("API-KEY")
		if _, exists := apiKeys[apiKey]; !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
			log.Panicf("Unauthorized Access")
			c.Abort()
			return
		}
		c.Next()
	}
}
