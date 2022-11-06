package main

import (
	"fmt"
	"github.com/itachilee/ginblog/internal/models"
	"github.com/itachilee/ginblog/internal/redis"
	"github.com/itachilee/ginblog/internal/routers"
	"github.com/itachilee/ginblog/internal/setting"
	"net/http"
	// "sync"
)

func init() {

	setting.Setup()
	models.InitDb()
	redis.InitRedis()
}

func main() {

	// endless.DefaultReadTimeOut = setting.ServerSetting.ReadTimeout
	// endless.DefaultWriteTimeOut = setting.ServerSetting.WriteTimeout
	// endless.DefaultMaxHeaderBytes = 1 << 20
	// endPoint := fmt.Sprintf(":%d", setting.ServerSetting.HttpPort)
	// server := endless.NewServer(endPoint, routers.InitRouter())
	// server.BeforeBegin = func(add string) {
	// 	log.Printf("Actual pid is %d", syscall.Getpid())
	// }

	// err := server.ListenAndServe()
	// if err != nil {
	// 	log.Printf("Server err:%v", err)
	// }
	// util.InitCron()

	routersInit := routers.InitRouter()

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.ServerSetting.HttpPort),
		Handler:        routersInit,
		ReadTimeout:    setting.ServerSetting.ReadTimeout,
		WriteTimeout:   setting.ServerSetting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()

}
