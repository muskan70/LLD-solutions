package score

import (
	"bufio"
	"context"
	"intuitMc/domain/user"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func ReadScoreFromFile(ctx context.Context) error {
	filename := ScoreFileName + time.Now().Add(-1*time.Minute).Format("2006-01-02 15:04")
	readFile, err := os.Open(filename)
	if err != nil {
		log.Println(err)
		return err
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		scoreRecord := strings.Split(fileScanner.Text(), "|")
		userId, _ := strconv.Atoi(scoreRecord[0])
		score, _ := strconv.Atoi(scoreRecord[1])
		if err := user.UpdateUserScores(ctx, userId, score); err != nil {
			log.Println(err)
			continue
		}
	}

	readFile.Close()
	return nil
}
