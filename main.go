package main

import (
	"fmt"
	"sync"
	"time"
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

	number := 10
	g := New(10)
	for i := 0; i < number; i++ {
		wg.Add(1)
		value := i
		goFunc := func() {
			// 做一些业务逻辑处理
			fmt.Printf("go func: %d\n", value)
			time.Sleep(time.Second)
			wg.Done()
		}
		g.Run(goFunc)
	}
	wg.Wait()
}
