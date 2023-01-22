package middleware

import "github.com/gin-gonic/gin"

func AppInfo() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("app_name", "leonlee's blog service")
		c.Set("app_version", "v1.0.0")
		c.Next()
	}
}
