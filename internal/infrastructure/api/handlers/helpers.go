package handlers

import (
	"github.com/gin-gonic/gin"
)

// Helper function to add security headers for successful responses
func addSuccessHeaders(c *gin.Context) {
	// Indicates the response was successful and validated
	c.Header("X-Request-ID", c.GetString("request_id"))

	// Prevent response caching for sensitive data
	c.Header("Cache-Control", "no-cache, no-store, must-revalidate, max-age=0")
}

// Helper function to add security headers for error responses
func addErrorHeaders(c *gin.Context) {
	// Indicates this is an error response
	c.Header("X-Error-Response", "true")

	// Prevent caching of error pages
	c.Header("Cache-Control", "no-cache, no-store, must-revalidate, max-age=0")

	// Request ID for error tracking
	c.Header("X-Request-ID", c.GetString("request_id"))
}

// Helper function to add validation headers
func addValidationHeaders(c *gin.Context) {
	// Indicates request was validated
	c.Header("X-Validated", "true")

	// Content security policy for JSON responses
	c.Header("Content-Type", "application/json; charset=utf-8")
}
