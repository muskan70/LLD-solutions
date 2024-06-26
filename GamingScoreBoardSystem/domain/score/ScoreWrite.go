package score

import (
	"context"
	"fmt"
	"intuitMc/requests"
	"log"
	"os"
)

const ScoreFileName = "score.txt"

func PushScoreToFile(ctx context.Context, req *requests.PushScoreRequest) error {
	f, err := os.OpenFile(ScoreFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	newLine := fmt.Sprintf("%d|%d", req.UserId, req.Score)
	_, err = fmt.Fprintln(f, newLine)
	if err != nil {
		log.Println(err.Error())
		f.Close()
		return err
	}
	err = f.Close()
	if err != nil {
		log.Println(err.Error())
		return err
	}
	return nil
}
