package middleware

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

// LoggingMiddleware logs the details of each HTTP request, including method, path, status, and duration.
func LoggingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Record the start time of the request
		startTime := time.Now()

		// Process the request
		c.Next()

		// Calculate the duration taken to handle the request
		duration := time.Since(startTime)

		// Log the request details
		fmt.Printf(
			"Request Method: %s | Path: %s | Status: %d | Duration: %v\n",
			c.Request.Method,
			c.Request.URL.Path,
			c.Writer.Status(),
			duration,
		)
	}
}
