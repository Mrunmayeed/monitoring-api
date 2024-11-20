package main

import (
	"log"
	"monitoring-api/handlers"
	"monitoring-api/middleware"
	"os"

	"github.com/gin-gonic/gin"
)

func setup() *os.File {

	// Open the log file
	f, err := os.OpenFile("temp/log/monitoring-api.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Error opening log file: %v", err)
	}

	// Set the logfile as default logging output
	log.SetOutput(f)
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	// Set the logfile as default logging output for gin server
	gin.DefaultWriter = f
	gin.DefaultErrorWriter = f

	// Remove debug logs
	gin.SetMode(gin.ReleaseMode)
	return f
}

func main() {

	f := setup()
	defer f.Close()

	// Initialize the Gin router
	log.Printf("Starting server and listening on port:8080")
	router := gin.Default()

	// Apply API key authentication middleware
	router.Use(middleware.ApiKeyAuthMiddleware())

	// Routes
	router.GET("/cpu", handlers.GetCPUUsage)
	router.GET("/memory", handlers.GetMemoryUsage)
	router.GET("/disk", handlers.GetDiskUsage)
	router.GET("/bandwidth", handlers.GetBandwidthUsage)

	// Start the server
	if err := router.Run("0.0.0.0:8080"); err != nil {
		log.Fatalf("Error starting the server: %v", err)
	}
}
