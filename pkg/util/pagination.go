package util

import (
	"github.com/gin-gonic/gin"
	"github.com/itachilee/ginblog/global"
	"github.com/unknwon/com"
)

func GetPage(c *gin.Context) int {
	result := 0
	page, _ := com.StrTo(c.Query("page")).Int()
	if page > 0 {
		result = (page - 1) * global.AppSetting.PageSize
	}
	return result
}
