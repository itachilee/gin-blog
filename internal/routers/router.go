package routers

import (
	"github.com/itachilee/ginblog/internal/middleware/jwt"
	api2 "github.com/itachilee/ginblog/internal/routers/api"
	v12 "github.com/itachilee/ginblog/internal/routers/api/v1"
	"github.com/itachilee/ginblog/internal/setting"
	"github.com/itachilee/ginblog/internal/upload"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	gin.SetMode(setting.ServerSetting.RunMode)
	r.StaticFS("/upload/images", http.Dir(upload.GetImageFullPath()))
	// r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.GET("/auth", api2.GetAuth)
	r.POST("/upload", api2.UploadImage)
	apiv1 := r.Group("/api/v1")
	apiv1.Use(jwt.JWT())
	{
		//获取标签列表
		apiv1.GET("/tags", v12.GetTags)
		//新建标签
		apiv1.POST("/tags", v12.AddTag)
		//更新指定标签
		apiv1.PUT("/tags/:id", v12.EditTag)
		//删除指定标签
		apiv1.DELETE("/tags/:id", v12.DeleteTag)

		//获取文章列表
		apiv1.GET("/articles", v12.GetArticles)
		//获取指定文章
		apiv1.GET("/articles/:id", v12.GetArticle)
		//新建文章
		apiv1.POST("/articles", v12.AddArticle)
		//更新指定文章
		apiv1.PUT("/articles/:id", v12.EditArticle)
		//删除指定文章
		apiv1.DELETE("/articles/:id", v12.DeleteArticle)
	}
	return r
}
