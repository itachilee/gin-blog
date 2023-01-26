package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/itachilee/ginblog/global"
	"github.com/itachilee/ginblog/internal/model"
	"github.com/itachilee/ginblog/internal/routers"
	"github.com/itachilee/ginblog/pkg/tracer"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	// "sync"
)

func init() {

	global.Setup()

	setupDBEngine()
	global.DBEngine.AutoMigrate(&model.Article{}, &model.Tag{}, &model.ArticleTag{}, &model.Auth{})
	setupTracer()
	setUpFlag()
	//redis.InitRedis()
}
func setupDBEngine() error {

	var err error
	global.DBEngine, err = model.NewDBEngine()
	if err != nil {
		return err
	}
	return nil
}

func setupTracer() error {
	jaegerTracer, _, err := tracer.NewJaegerTracer(
		"blog-service",
		"127.0.0.1:6831",
	)

	if err != nil {
		return err
	}
	global.Tracer = jaegerTracer
	return nil
}

var (
	port    string
	runMode string
	config  string
)

func setUpFlag() error {
	flag.StringVar(&port, "port", "", "启动端口")
	flag.StringVar(&runMode, "mode", "", "启动模式")
	flag.StringVar(&config, "config", "configs/", "指定要使用的配置文件路径")
	flag.Parse()
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

	go func() {
		err := s.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			fmt.Printf("s.ListenAndServe() err: %v \n", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	fmt.Printf("shuting down server \n")
	ctx, cancer := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancer()
	if err := s.Shutdown(ctx); err != nil {
		fmt.Errorf("server forced to shutdown: %v \n", err)
	}
	fmt.Printf("server shutdown")
}
