package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/itachilee/ginblog/pkg/casbin"
	"log"
	"net/http"
)

func CasbinHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 获取请求的URI
		obj := ctx.Request.URL.RequestURI()
		// 获取请求方法
		act := ctx.Request.Method
		// 获取用户的角色
		sub := "admin"
		e := casbin.InitEnforce()
		fmt.Println(obj, act, sub)
		// 判断策略中是否存在
		success, err := e.Enforce(sub, obj, act)
		if err != nil {
			log.Println("权限验证异常")
			ctx.Abort()
		}
		if success {
			log.Println("恭喜您,权限验证通过")
			ctx.Next()
		} else {
			log.Printf("errcode.Enforce err: %s", "很遗憾,权限验证没有通过")
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"msg": "很遗憾,权限验证没有通过",
			})
			ctx.Abort()
			return
		}
	}
}
