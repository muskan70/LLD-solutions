package score

import (
	"bufio"
	"context"
	"intuitMc/domain/user"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"
)

func ReadScoreFromFile(ctx context.Context) error {
	readFile, err := os.Open(ScoreFileName)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var wg sync.WaitGroup
	for fileScanner.Scan() {
		scoreRecord := strings.Split(fileScanner.Text(), "|")
		userId, _ := strconv.Atoi(scoreRecord[0])
		score, _ := strconv.Atoi(scoreRecord[1])
		wg.Add(1)
		go func() {
			defer wg.Done()
			if err := user.UpdateUserScores(ctx, userId, score); err != nil {
				log.Println(err.Error())
			}
		}()
	}
	readFile.Close()
	wg.Wait()
	return nil
}
