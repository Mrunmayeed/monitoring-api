package handlers

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/net"
	"github.com/stretchr/testify/assert"
)

func TestGetCPUUsage_Success(t *testing.T) {
	// Mock function for fetching CPU usage
	mockFetchCPUUsage := func(iinterval time.Duration, percpu bool) ([]float64, error) {
		return []float64{42.5}, nil
	}

	// Set up Gin router
	router := gin.Default()
	router.GET("/cpu", GetCPUUsage(mockFetchCPUUsage))

	// Create a test request
	req, _ := http.NewRequest("GET", "/cpu", nil)
	w := httptest.NewRecorder()

	// Send the request
	router.ServeHTTP(w, req)

	// Assert the response
	assert.Equal(t, http.StatusOK, w.Code)
	assert.JSONEq(t, `{"cpu_usage_percent": 42.5}`, w.Body.String())
}

func TestGetCPUUsage_Failure(t *testing.T) {
	// Mock function to simulate an error
	mockFetchCPUUsage := func(interval time.Duration, percpu bool) ([]float64, error) {
		return nil, errors.New("failed to fetch CPU usage")
	}

	// Set up Gin router
	router := gin.Default()
	router.GET("/cpu", GetCPUUsage(mockFetchCPUUsage))

	// Create a test request
	req, _ := http.NewRequest("GET", "/cpu", nil)
	w := httptest.NewRecorder()

	// Send the request
	router.ServeHTTP(w, req)

	// Assert the response
	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.JSONEq(t, `{"error": "Error fetching CPU usage"}`, w.Body.String())
}

func TestGetMemoryUsage_Success(t *testing.T) {
	// Mock implementation of MemoryUsageFetcher
	mockFetchMemoryUsage := func() (*mem.VirtualMemoryStat, error) {
		return &mem.VirtualMemoryStat{
			Total:       16000000000, // 16 GB
			Used:        8000000000,  // 8 GB
			UsedPercent: 50.0,        // 50% usage
		}, nil
	}

	// Set up Gin router
	router := gin.Default()
	router.GET("/memory", GetMemoryUsage(mockFetchMemoryUsage))

	// Simulate an HTTP GET request
	req, _ := http.NewRequest("GET", "/memory", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Assert the response
	assert.Equal(t, http.StatusOK, w.Code)
	assert.JSONEq(t, `{
		"total_memory": 16000000000,
		"used_memory": 8000000000,
		"memory_percent": 50.0
	}`, w.Body.String())
}

func TestGetMemoryUsage_Failure(t *testing.T) {
	// Mock implementation to simulate an error
	mockFetchMemoryUsage := func() (*mem.VirtualMemoryStat, error) {
		return nil, errors.New("failed to fetch memory usage")
	}

	// Set up Gin router
	router := gin.Default()
	router.GET("/memory", GetMemoryUsage(mockFetchMemoryUsage))

	// Simulate an HTTP GET request
	req, _ := http.NewRequest("GET", "/memory", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Assert the response
	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.JSONEq(t, `{"error": "Error fetching memory usage"}`, w.Body.String())
}

func TestGetDiskUsage_Success(t *testing.T) {
	// Mock implementation of DiskUsageFetcher
	mockFetchDiskUsage := func(path string) (*disk.UsageStat, error) {
		return &disk.UsageStat{
			Total:       100000000000, // 100 GB
			Used:        50000000000,  // 50 GB
			UsedPercent: 50.0,         // 50% usage
		}, nil
	}

	// Set up Gin router
	router := gin.Default()
	router.GET("/disk", GetDiskUsage(mockFetchDiskUsage))

	// Simulate an HTTP GET request
	req, _ := http.NewRequest("GET", "/disk", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Assert the response
	assert.Equal(t, http.StatusOK, w.Code)
	assert.JSONEq(t, `{
		"total_disk": 100000000000,
		"used_disk": 50000000000,
		"disk_percent": 50.0
	}`, w.Body.String())
}

func TestGetDiskUsage_Failure(t *testing.T) {
	// Mock implementation to simulate an error
	mockFetchDiskUsage := func(path string) (*disk.UsageStat, error) {
		return nil, errors.New("failed to fetch disk usage")
	}

	// Set up Gin router
	router := gin.Default()
	router.GET("/disk", GetDiskUsage(mockFetchDiskUsage))

	// Simulate an HTTP GET request
	req, _ := http.NewRequest("GET", "/disk", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Assert the response
	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.JSONEq(t, `{"error": "Error fetching disk usage"}`, w.Body.String())
}

func TestGetBandwidthUsage_Success(t *testing.T) {
	// Mock implementation of BandwidthUsageFetcher
	mockFetchBandwidthUsage := func(pernic bool) ([]net.IOCountersStat, error) {
		return []net.IOCountersStat{
			{
				BytesSent: 1000000000, // 1 GB
				BytesRecv: 500000000,  // 500 MB
			},
		}, nil
	}

	// Set up Gin router
	router := gin.Default()
	router.GET("/bandwidth", GetBandwidthUsage(mockFetchBandwidthUsage))

	// Simulate an HTTP GET request
	req, _ := http.NewRequest("GET", "/bandwidth", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Assert the response
	assert.Equal(t, http.StatusOK, w.Code)
	assert.JSONEq(t, `{
		"bytes_sent": 1000000000,
		"bytes_received": 500000000
	}`, w.Body.String())
}

func TestGetBandwidthUsage_Failure(t *testing.T) {
	// Mock implementation to simulate an error
	mockFetchBandwidthUsage := func(pernic bool) ([]net.IOCountersStat, error) {
		return nil, errors.New("failed to fetch network usage")
	}

	// Set up Gin router
	router := gin.Default()
	router.GET("/bandwidth", GetBandwidthUsage(mockFetchBandwidthUsage))

	// Simulate an HTTP GET request
	req, _ := http.NewRequest("GET", "/bandwidth", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Assert the response
	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.JSONEq(t, `{"error": "Error fetching bandwidth usage"}`, w.Body.String())
}
