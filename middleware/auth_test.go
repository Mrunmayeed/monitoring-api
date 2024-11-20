package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestApiKeyAuthMiddleware_ValidKey(t *testing.T) {
	// Mock API keys
	mockApiKeys := map[string]bool{
		"valid-key": true,
	}

	// Set up Gin router
	router := gin.Default()
	router.Use(ApiKeyAuthMiddleware(mockApiKeys))
	router.GET("/protected", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Access granted"})
	})

	// Simulate a request with a valid API key
	req, _ := http.NewRequest("GET", "/protected", nil)
	req.Header.Set("API-KEY", "valid-key")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Assert the response
	assert.Equal(t, http.StatusOK, w.Code)
	assert.JSONEq(t, `{"message": "Access granted"}`, w.Body.String())
}

func TestApiKeyAuthMiddleware_InvalidKey(t *testing.T) {
	// Mock API keys
	mockApiKeys := map[string]bool{
		"valid-key": true,
	}

	// Set up Gin router
	router := gin.Default()
	router.Use(ApiKeyAuthMiddleware(mockApiKeys))
	router.GET("/protected", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Access granted"})
	})

	// Simulate a request with an invalid API key
	req, _ := http.NewRequest("GET", "/protected", nil)
	req.Header.Set("API-KEY", "invalid-key")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Assert the response
	assert.Equal(t, http.StatusUnauthorized, w.Code)
	assert.JSONEq(t, `{"message": "Unauthorized"}`, w.Body.String())
}

func TestApiKeyAuthMiddleware_NoKey(t *testing.T) {
	// Mock API keys
	mockApiKeys := map[string]bool{
		"valid-key": true,
	}

	// Set up Gin router
	router := gin.Default()
	router.Use(ApiKeyAuthMiddleware(mockApiKeys))
	router.GET("/protected", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Access granted"})
	})

	// Simulate a request without an API key
	req, _ := http.NewRequest("GET", "/protected", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Assert the response
	assert.Equal(t, http.StatusUnauthorized, w.Code)
	assert.JSONEq(t, `{"message": "Unauthorized"}`, w.Body.String())
}
