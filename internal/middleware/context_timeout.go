package middleware

import (
	"context"
	"github.com/gin-gonic/gin"
	"time"
)

// Middleware 超时控制
func ContextTimeout(t time.Duration) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(c.Request.Context(), t)
		defer cancel()
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}
