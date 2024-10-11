package util

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(next gin.HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		fmt.Println("Authorization header:", authHeader)

		if authHeader == "" {
			fmt.Println("Authorization header is missing")
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is missing"})
			c.Abort()
			return
		}

		// Check if the Authorization header has the "Bearer " prefix
		const bearerPrefix = "Bearer "
		if len(authHeader) > len(bearerPrefix) && authHeader[:len(bearerPrefix)] == bearerPrefix {
			tokenString := authHeader[len(bearerPrefix):]
			fmt.Println("Token string:", tokenString)
			// Validate the token
			_, err := ValidateToken(tokenString)
			if err != nil {
				fmt.Println("Invalid token:", err)
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token", "details": err.Error()})
				c.Abort()
				return
			}
			// Set the token in the context
			c.Set("token", tokenString)
			fmt.Println("Token set in context:", tokenString)
			next(c) // Call the next handler
		} else {
			fmt.Println("Invalid Authorization header format")
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Authorization header format"})
			c.Abort()
			return
		}
	}
}
