package util

import (
	models2 "github.com/itachilee/ginblog/internal/models"
	"log"
	"time"

	"github.com/robfig/cron"
)

func InitCron() {
	log.Println("Starting ...")

	c := cron.New()

	c.AddFunc("* 1 * * * *", func() {
		log.Println("Running models.CleanAllTag...")
		models2.CleanAllTag()
	})
	c.AddFunc("* 1 * * * *", func() {
		log.Println("Running models.CleanAllArticle...")
		models2.CleanAllArticle()
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
