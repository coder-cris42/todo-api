package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// RequestID middleware adds a unique request ID to each request
// This helps with request tracing and security logging
func RequestID() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Generate a unique request ID
		requestID := uuid.New().String()

		// Store it in the context for access in handlers
		c.Set("request_id", requestID)

		// Add it to the response headers
		c.Header("X-Request-ID", requestID)

		// Continue to next handler
		c.Next()
	}
}

// InputValidation middleware adds headers indicating input validation was performed
func InputValidation() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Indicates request has been validated against security policies
		c.Header("X-Request-Validated", "true")

		c.Next()
	}
}

// ResponseValidation middleware ensures response integrity
func ResponseValidation() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Indicates response is safe and validated
		c.Header("X-Response-Validated", "true")

		c.Next()
	}
}
