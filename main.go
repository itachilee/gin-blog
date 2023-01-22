package main

import (
	"fmt"
	"github.com/itachilee/ginblog/global"
	"github.com/itachilee/ginblog/internal/model"
	"github.com/itachilee/ginblog/internal/routers"
	"net/http"
	// "sync"
)

func init() {

	global.Setup()

	SetupDBEngine()
	global.DBEngine.AutoMigrate(&model.Article{}, &model.Tag{}, &model.ArticleTag{}, &model.Auth{})

	//redis.InitRedis()
}
func SetupDBEngine() error {

	var err error
	global.DBEngine, err = model.NewDBEngine()
	if err != nil {
		return err
	}
	return nil
}

// 添加注释以描述 server 信息
// @title           Swagger Example API
// @version         1.0
// @description     This is a sample server celler server.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/v1

// @securityDefinitions.basic  BasicAuth
func main() {

	routersInit := routers.InitRouter()

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", global.ServerSetting.HttpPort),
		Handler:        routersInit,
		ReadTimeout:    global.ServerSetting.ReadTimeout,
		WriteTimeout:   global.ServerSetting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()

}
