package middleware

import (
	"database/sql"
	"encoding/base64"
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
)

// AuthMiddleware validates Basic Authentication credentials against the database
func AuthMiddleware(db *sql.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Retrieve the Authorization header
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Basic ") {
			fmt.Println("Missing or invalid Authorization header")
			ctx.JSON(401, gin.H{"error": "Unauthorized"})
			ctx.Abort()
			return
		}

		// Decode the Base64-encoded credentials
		decoded, err := base64.StdEncoding.DecodeString(strings.TrimPrefix(authHeader, "Basic "))
		if err != nil {
			fmt.Println("Failed to decode Authorization header:", err)
			ctx.JSON(401, gin.H{"error": "Invalid Authorization Header"})
			ctx.Abort()
			return
		}

		// Extract username and password
		credentials := strings.SplitN(string(decoded), ":", 2)
		if len(credentials) != 2 {
			fmt.Println("Invalid credentials format:", string(decoded))
			ctx.JSON(401, gin.H{"error": "Invalid Credentials"})
			ctx.Abort()
			return
		}

		username, password := credentials[0], credentials[1]

		// Validate credentials against the database
		var storedPassword string
		query := "SELECT password FROM users WHERE username = ?"
		err = db.QueryRow(query, username).Scan(&storedPassword)
		if err != nil {
			fmt.Println("User not found or error querying database:", err)
			ctx.JSON(401, gin.H{"error": "Unauthorized"})
			ctx.Abort()
			return
		}

		// Check if passwords match
		if storedPassword != password {
			fmt.Println("Password mismatch")
			ctx.JSON(401, gin.H{"error": "Unauthorized"})
			ctx.Abort()
			return
		}

		fmt.Println("Authentication successful")
		ctx.Next()
	}
}
