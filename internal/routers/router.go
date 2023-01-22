package routers

import (
	"github.com/itachilee/ginblog/global"
	"github.com/itachilee/ginblog/internal/middleware"
	v1 "github.com/itachilee/ginblog/internal/routers/api/v1"
	"github.com/itachilee/ginblog/pkg/limiter"
	"github.com/itachilee/ginblog/pkg/upload"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/itachilee/ginblog/docs" // main 文件中导入 docs 包

	"github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitRouter() *gin.Engine {
	var r *gin.Engine

	if global.ServerSetting.RunMode == "debug" {
		r = gin.Default()
		r.Use(gin.Logger())
		r.Use(gin.Recovery())
	} else {
		r = gin.New()
		r.Use(middleware.LoggerToFile())
		r.Use(gin.Recovery())
	}
	r.Use(middleware.AppInfo())
	r.Use(middleware.Translations())
	r.Use(middleware.ContextTimeout(global.AppSetting.ContextTimeout))

	var methodLimiters = limiter.NewMethodLimiter().AddBuckets(
		limiter.LimiterBucketRule{
			Key:          "/auth",
			FillInterval: time.Second,
			Capacity:     10,
			Quantum:      10,
		})
	r.Use(middleware.RateLimiter(methodLimiters))
	gin.SetMode(global.ServerSetting.RunMode)
	r.StaticFS("/upload/images", http.Dir(upload.GetImageFullPath()))
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.POST("/auth", v1.GetAuth)

	article := v1.Article{}
	tag := v1.Tag{}
	apiv1 := r.Group("/api/v1")
	//apiv1.Use(middleware.JWT())
	{
		apiv1.Use(middleware.JWT())

		apiv1.GET("/hello", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"status":  http.StatusOK,
				"message": "pong",
			})
		})

		apiv1.POST("/tags", tag.Create)
		apiv1.DELETE("/tags/:id", tag.Delete)
		apiv1.PUT("/tags", tag.Update)
		apiv1.GET("/tags", tag.List)
		apiv1.GET("/tags/:id", tag.Get)

		apiv1.POST("/articles", article.Create)
		apiv1.DELETE("/articles/:id", article.Delete)
		apiv1.PUT("/articles", article.Update)
		apiv1.GET("/articles", article.List)
		apiv1.GET("/articles/:id", article.Get)
	}
	return r
}
