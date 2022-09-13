package cron

import (
	"context"
	"intuitMc/domain/score"
	"log"

	"github.com/robfig/cron/v3"
)

func ScoreProcessing() {
	c := cron.New()
	c.AddFunc("@every 1m", func() {
		ctx := context.Background()
		err := score.ReadScoreFromFile(ctx)
		if err != nil {
			log.Println(err)
			return
		}
	})
	c.Start()
}
