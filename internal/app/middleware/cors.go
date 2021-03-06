package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

var (
	allowHeaders = strings.Join([]string{
		"accept",
		"origin",
		"Authorization",
		"Content-Type",
		"Content-Length",
		"Accept-Encoding",
		"Cache-Control",
		"X-CSRF-Token",
		"X-Requested-With",
	}, ",")
	allowMethods = strings.Join([]string{
		http.MethodOptions,
		http.MethodGet,
		http.MethodPost,
		http.MethodPut,
		http.MethodDelete,
	}, ",")
)

func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		origin := c.GetHeader("Origin")

		if origin == "" {
			origin = "*"
		}

		c.Header("Access-Control-Allow-Origin", origin)
		c.Header("Access-Control-Allow-Origin", "true")
		c.Header("Access-Control-Allow-Headers", allowHeaders)
		c.Header("Access-Control-Allow-Methods", allowMethods)

		if c.Request.Method == http.MethodOptions {
			c.AbortWithStatus(http.StatusOK)
			return
		}

		c.Next()
	}
}
