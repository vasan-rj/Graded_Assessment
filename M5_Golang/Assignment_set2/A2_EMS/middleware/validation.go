package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// ValidationMiddleware ensures that incoming requests have the correct Content-Type.
func ValidationMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Check if the Content-Type header is set to "application/json"
		if c.Request.Header.Get("Content-Type") != "application/json" {
			// Respond with a 400 Bad Request status and an error message
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid content type"})
			c.Abort()
			return
		}

		// Proceed with the next middleware or handler if the Content-Type is valid
		c.Next()
	}
}
