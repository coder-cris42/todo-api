package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"todo-api/internal/infrastructure/api/handlers"
	"todo-api/internal/infrastructure/database/repositories"
	"todo-api/internal/middleware"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// TestSecurityHeadersGlobal tests that global security headers are present in all responses
func TestSecurityHeadersGlobal(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.New()

	// Apply security middlewares
	router.Use(middleware.SecurityHeaders())
	router.Use(middleware.CORSHeaders())
	router.Use(middleware.RequestID())

	// Simple test endpoint
	router.GET("/test", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "test"})
	})

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/test", nil)

	router.ServeHTTP(w, req)

	// Assert global security headers
	assert.Equal(t, "nosniff", w.Header().Get("X-Content-Type-Options"))
	assert.Equal(t, "DENY", w.Header().Get("X-Frame-Options"))
	assert.Equal(t, "1; mode=block", w.Header().Get("X-XSS-Protection"))
	assert.Equal(t, "strict-origin-when-cross-origin", w.Header().Get("Referrer-Policy"))
	assert.NotEmpty(t, w.Header().Get("Strict-Transport-Security"))
	assert.NotEmpty(t, w.Header().Get("Content-Security-Policy"))
	assert.NotEmpty(t, w.Header().Get("Permissions-Policy"))
	assert.NotEmpty(t, w.Header().Get("Cache-Control"))
	assert.Equal(t, "no-cache", w.Header().Get("Pragma"))
	assert.Equal(t, "0", w.Header().Get("Expires"))
	assert.Equal(t, "1.0", w.Header().Get("X-API-Version"))
	assert.Equal(t, "Todo-API", w.Header().Get("X-Powered-By"))
	assert.NotEmpty(t, w.Header().Get("X-Request-ID"))
}

// TestCORSHeaders tests CORS headers are present
func TestCORSHeaders(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.New()

	router.Use(middleware.CORSHeaders())

	router.GET("/test", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "test"})
	})

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/test", nil)
	req.Header.Set("Origin", "http://example.com")

	router.ServeHTTP(w, req)

	// Assert CORS headers
	assert.NotEmpty(t, w.Header().Get("Access-Control-Allow-Origin"))
	assert.NotEmpty(t, w.Header().Get("Access-Control-Allow-Methods"))
	assert.NotEmpty(t, w.Header().Get("Access-Control-Allow-Headers"))
	assert.NotEmpty(t, w.Header().Get("Access-Control-Expose-Headers"))
}

// TestCORSPreflight tests preflight OPTIONS request handling
func TestCORSPreflight(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.New()

	router.Use(middleware.CORSHeaders())

	router.OPTIONS("/test", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "ok"})
	})

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("OPTIONS", "/test", nil)
	req.Header.Set("Origin", "http://example.com")

	router.ServeHTTP(w, req)

	// Preflight should return 204
	assert.Equal(t, http.StatusNoContent, w.Code)
}

// TestRequestIDGeneration tests that X-Request-ID is generated
func TestRequestIDGeneration(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.New()

	router.Use(middleware.RequestID())

	router.GET("/test", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "test"})
	})

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/test", nil)

	router.ServeHTTP(w, req)

	requestID := w.Header().Get("X-Request-ID")
	assert.NotEmpty(t, requestID)
	// UUID format should be like: 123e4567-e89b-12d3-a456-426614174000
	assert.Len(t, requestID, 36) // UUID length
}

// TestHandlerSecurityHeaders tests that handlers add security headers
func TestHandlerSecurityHeaders(t *testing.T) {
	gin.SetMode(gin.TestMode)

	// Mock repository
	mockRepo := &repositories.TaskRepository{}
	handler := handlers.NewTaskHandler(mockRepo)

	router := gin.New()
	router.Use(middleware.SecurityHeaders())
	router.POST("/todo", handler.CreateTask)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/todo", nil)
	req.Header.Set("Content-Type", "application/json")

	router.ServeHTTP(w, req)

	// Even on error, security headers should be present
	assert.Equal(t, "nosniff", w.Header().Get("X-Content-Type-Options"))
	assert.Equal(t, "DENY", w.Header().Get("X-Frame-Options"))
	assert.Equal(t, "application/json; charset=utf-8", w.Header().Get("Content-Type"))
}

// TestCachePreventionHeaders tests that cache prevention headers are set correctly
func TestCachePreventionHeaders(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.New()

	router.Use(middleware.SecurityHeaders())

	router.GET("/sensitive", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "sensitive"})
	})

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/sensitive", nil)

	router.ServeHTTP(w, req)

	// Verify cache prevention
	assert.Contains(t, w.Header().Get("Cache-Control"), "no-cache")
	assert.Contains(t, w.Header().Get("Cache-Control"), "no-store")
	assert.Contains(t, w.Header().Get("Cache-Control"), "must-revalidate")
	assert.Equal(t, "0", w.Header().Get("Expires"))
	assert.Equal(t, "no-cache", w.Header().Get("Pragma"))
}

// TestInputValidationHeader tests that input validation header is present
func TestInputValidationHeader(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.New()

	router.Use(middleware.InputValidation())

	router.GET("/test", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "test"})
	})

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/test", nil)

	router.ServeHTTP(w, req)

	assert.Equal(t, "true", w.Header().Get("X-Request-Validated"))
}

// TestResponseValidationHeader tests that response validation header is present
func TestResponseValidationHeader(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.New()

	router.Use(middleware.ResponseValidation())

	router.GET("/test", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "test"})
	})

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/test", nil)

	router.ServeHTTP(w, req)

	assert.Equal(t, "true", w.Header().Get("X-Response-Validated"))
}

// TestHTST tests Strict-Transport-Security header
func TestHSTS(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.New()

	router.Use(middleware.SecurityHeaders())

	router.GET("/test", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "test"})
	})

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/test", nil)

	router.ServeHTTP(w, req)

	hsts := w.Header().Get("Strict-Transport-Security")
	assert.NotEmpty(t, hsts)
	assert.Contains(t, hsts, "max-age=31536000")
	assert.Contains(t, hsts, "includeSubDomains")
	assert.Contains(t, hsts, "preload")
}

// TestCSP tests Content-Security-Policy header
func TestCSP(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.New()

	router.Use(middleware.SecurityHeaders())

	router.GET("/test", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "test"})
	})

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/test", nil)

	router.ServeHTTP(w, req)

	csp := w.Header().Get("Content-Security-Policy")
	assert.NotEmpty(t, csp)
	assert.Contains(t, csp, "default-src 'self'")
	assert.Contains(t, csp, "script-src 'self'")
	assert.Contains(t, csp, "frame-ancestors 'none'")
}

// TestPermissionsPolicy tests Permissions-Policy header
func TestPermissionsPolicy(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.New()

	router.Use(middleware.SecurityHeaders())

	router.GET("/test", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "test"})
	})

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/test", nil)

	router.ServeHTTP(w, req)

	permPolicy := w.Header().Get("Permissions-Policy")
	assert.NotEmpty(t, permPolicy)
	assert.Contains(t, permPolicy, "geolocation=()")
	assert.Contains(t, permPolicy, "microphone=()")
	assert.Contains(t, permPolicy, "camera=()")
}
