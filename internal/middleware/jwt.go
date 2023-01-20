package middleware

import (
	e2 "github.com/itachilee/ginblog/pkg/errcode"
	"github.com/itachilee/ginblog/pkg/util"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		var data interface{}
		code = e2.SUCCESS
		token := c.Query("token")
		if token == "" {
			code = e2.INVALID_PARAMS
		} else {
			claims, err := util.ParseToken(token)
			if err != nil {
				code = e2.ERROR_AUTH_CHECK_TOKEN_FAIL
			} else if time.Now().Unix() > claims.ExpiresAt {
				code = e2.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
			}
		}
		if code != e2.SUCCESS {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": code,
				"msg":  e2.GetMsg(code),
				"data": data,
			})
			c.Abort()
			return
		}
	}
}
