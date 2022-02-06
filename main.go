package main

import (
	"collyD/pkg/setting"
	"collyD/pkg/util"
	"collyD/routers"
	"fmt"
	"log"
	"sync"
	"syscall"

	"github.com/fvbock/endless"
	// "sync"
)

type Glimit struct {
	n int
	c chan struct{}
}

func New(n int) *Glimit {
	return &Glimit{
		n: n,
		c: make(chan struct{}, n),
	}
}

func (g *Glimit) Run(f func()) {
	g.c <- struct{}{}
	go func() {
		f()
		<-g.c
	}()

}

var wg = sync.WaitGroup{}

func main() {

	endless.DefaultReadTimeOut = setting.ReadTimeout
	endless.DefaultWriteTimeOut = setting.WriteTimeout
	endless.DefaultMaxHeaderBytes = 1 << 20
	endPoint := fmt.Sprintf(":%d", setting.HTTPPort)
	server := endless.NewServer(endPoint, routers.InitRouter())
	server.BeforeBegin = func(add string) {
		log.Printf("Actual pid is %d", syscall.Getpid())
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Printf("Server err:%v", err)
	}
	util.InitCron()

}
