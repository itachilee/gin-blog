package main

import (
	"collyD/models"
	"collyD/pkg/logging"
	"collyD/pkg/mqtt"
	"collyD/pkg/redis"
	"collyD/pkg/setting"
	"collyD/routers"
	"fmt"
	"net/http"
	// "sync"
)

func main() {

	setting.Setup()
	models.SetUp()
	logging.SetUp()
	redis.SetUp()
	go mqtt.SetUp()
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
