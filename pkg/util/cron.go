package util

import (
	"log"
	"time"

	"github.com/robfig/cron"
)

func InitCron() {
	log.Println("Starting ...")

	c := cron.New()

	c.AddFunc("* 1 * * * *", func() {
		log.Println("Running model.CleanAllTag...")
	})
	c.AddFunc("* 1 * * * *", func() {
		log.Println("Running model.CleanAllArticle...")
	})

	c.Start()

	t1 := time.NewTimer(time.Second * 10)
	for {
		select {
		case <-t1.C:
			t1.Reset(time.Second * 10)
		}
	}
}
