package handlers

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/net"
)

// GetCPUUsage returns the CPU usage percentage
type CPUUsageFetcher func(interval time.Duration, percpu bool) ([]float64, error)

func GetCPUUsage(fetchCPUUsage CPUUsageFetcher) gin.HandlerFunc {
	return func(c *gin.Context) {
		percentages, err := fetchCPUUsage(0, false)
		if err != nil {
			log.Println("Error fetching CPU usage:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching CPU usage"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"cpu_usage_percent": percentages[0]})
	}
}

// GetMemoryUsage returns the memory usage
type MemoryUsageFetcher func() (*mem.VirtualMemoryStat, error)

func GetMemoryUsage(fetchMemoryUsage MemoryUsageFetcher) gin.HandlerFunc {
	return func(c *gin.Context) {
		v, err := fetchMemoryUsage()
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
}

// GetDiskUsage returns the disk usage
type DiskUsageFetcher func(path string) (*disk.UsageStat, error)

func GetDiskUsage(fetchDiskUsage DiskUsageFetcher) gin.HandlerFunc {
	return func(c *gin.Context) {
		d, err := fetchDiskUsage("/")
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
}

// GetBandwidthUsage returns the network bandwidth usage
type BandwidthUsageFetcher func(pernic bool) ([]net.IOCountersStat, error)

func GetBandwidthUsage(fetchBandwidthUsage BandwidthUsageFetcher) gin.HandlerFunc {
	return func(c *gin.Context) {
		netIO, err := fetchBandwidthUsage(false)
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
}
