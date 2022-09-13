package score

import (
	"context"
	"fmt"
	"intuitMc/requests"
	"log"
	"os"
	"time"
)

const ScoreFileName = "score.txt"

func PushScoreToFile(ctx context.Context, req *requests.PushScoreRequest) error {
	filename := ScoreFileName + time.Now().Format("2006-01-02 15:04")
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
		return err
	}
	newLine := fmt.Sprintf("%d|%d", req.UserId, req.Score)
	_, err = fmt.Fprintln(f, newLine)
	if err != nil {
		log.Println(err)
		f.Close()
		return err
	}
	err = f.Close()
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
