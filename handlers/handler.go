package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/net"
)

// GetCPUUsage returns the CPU usage percentage
func GetCPUUsage(c *gin.Context) {
	percentages, err := cpu.Percent(0, false)
	if err != nil {
		log.Println("Error fetching CPU usage:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching CPU usage"})
		return
	}
	// log.Printf("200|  %s | CPU Usage returned | %s \n", c.Request.URL.Path, c.ClientIP())
	c.JSON(http.StatusOK, gin.H{"cpu_usage_percent": percentages[0]})
}

// GetMemoryUsage returns the memory usage
func GetMemoryUsage(c *gin.Context) {
	v, err := mem.VirtualMemory()
	if err != nil {
		log.Println("Error fetching memory usage:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching memory usage"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"total_memory":   v.Total,
		"used_memory":    v.Used,
		"memory_percent": v.UsedPercent,
	})
}

// GetDiskUsage returns the disk usage
func GetDiskUsage(c *gin.Context) {
	d, err := disk.Usage("/")
	if err != nil {
		log.Println("Error fetching disk usage:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching disk usage"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"total_disk":   d.Total,
		"used_disk":    d.Used,
		"disk_percent": d.UsedPercent,
	})
}

// GetBandwidthUsage returns the network bandwidth usage
func GetBandwidthUsage(c *gin.Context) {
	netIO, err := net.IOCounters(false)
	if err != nil {
		log.Println("Error fetching network usage:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching bandwidth usage"})
		return
	}

	// log.Printf("200 |  %s | Bandwidth Usage returned | %s \n", c.Request.URL.Path, c.Request.RequestURI)
	c.JSON(http.StatusOK, gin.H{
		"bytes_sent":     netIO[0].BytesSent,
		"bytes_received": netIO[0].BytesRecv,
	})
}
