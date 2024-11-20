package main

import (
	"log"
	"monitoring-api/handlers"
	"monitoring-api/middleware"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/net"
)

// Map of allowed API keys
var apiKeys = map[string]bool{
	"cs218secret": true,
}

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
	router.Use(middleware.ApiKeyAuthMiddleware(apiKeys))

	// Routes
	router.GET("/cpu", handlers.GetCPUUsage(cpu.Percent))
	router.GET("/memory", handlers.GetMemoryUsage(mem.VirtualMemory))
	router.GET("/disk", handlers.GetDiskUsage(disk.Usage))
	router.GET("/bandwidth", handlers.GetBandwidthUsage(net.IOCounters))

	// Start the server
	if err := router.Run("0.0.0.0:8080"); err != nil {
		log.Fatalf("Error starting the server: %v", err)
	}
}
