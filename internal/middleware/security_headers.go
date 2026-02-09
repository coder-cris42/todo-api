package middleware

import (
	"github.com/gin-gonic/gin"
)

// Securrity Headers according OWASP https://owasp.org/www-project-secure-headers/

func SecurityHeaders() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("X-Content-Type-Options", "nosniff")

		c.Header("X-Frame-Options", "DENY")

		c.Header("X-XSS-Protection", "1; mode=block")

		c.Header("Referrer-Policy", "strict-origin-when-cross-origin")

		c.Header("Strict-Transport-Security", "max-age=31536000; includeSubDomains; preload")

		c.Header("Content-Security-Policy", "default-src 'self'; script-src 'self'; style-src 'self' 'unsafe-inline'; img-src 'self' data: https:; font-src 'self'; connect-src 'self'; frame-ancestors 'none'; upgrade-insecure-requests")

		c.Header("Permissions-Policy", "geolocation=(), microphone=(), camera=(), payment=(), usb=(), accelerometer=(), gyroscope=(), magnetometer=(), vr=(), xr=(), ambient-light-sensor=()")

		c.Header("Cache-Control", "no-cache, no-store, must-revalidate, max-age=0")
		c.Header("Pragma", "no-cache")
		c.Header("Expires", "0")

		c.Header("Vary", "Accept-Encoding")

		c.Header("X-API-Version", "1.0")

		c.Header("X-Powered-By", "Todo-API")

		c.Header("Access-Control-Max-Age", "3600")

		c.Next()
	}
}

// Cross-Origin Resource Sharing headers
func CORSHeaders() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")

		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, PATCH, OPTIONS")

		c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization, Accept, Accept-Language, Content-Language")

		c.Header("Access-Control-Expose-Headers", "Content-Length, Content-Type, X-API-Version")

		c.Header("Access-Control-Allow-Credentials", "false")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
