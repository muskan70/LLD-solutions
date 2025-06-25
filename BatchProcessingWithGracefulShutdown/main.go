package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

type message struct {
	Val       int
	Timestamp time.Time
}

func produceData(ctx context.Context, dataBuffer chan message, wg *sync.WaitGroup) {
	defer wg.Done()
	defer close(dataBuffer)
	i := 0
	for {
		select {
		case <-ctx.Done():
			log.Println("Interrupt is detected")
			return
		default:
		}
		dataBuffer <- message{Val: i, Timestamp: time.Now()}
		i++
		randomN := rand.Intn(1000)
		time.Sleep(time.Duration(randomN) * time.Millisecond)
	}
}

func batchData(ctx context.Context, dataBuffer chan message, timeout time.Duration, fileSize int, wg *sync.WaitGroup) {
	defer wg.Done()
	var file []int
	lastTimeChecked := time.Now()
	for {
		select {
		case <-ctx.Done():
			for msg := range dataBuffer {
				file = append(file, msg.Val)
			}
			printFile(file)
			log.Println("Interrupt is detected")
			return
		case msg := <-dataBuffer:
			file = append(file, msg.Val)
			if len(file) == fileSize {
				fmt.Println(len(file))
				go printFile(file)
				file = make([]int, 0)
				lastTimeChecked = msg.Timestamp
			}
		default:
		}
		if time.Since(lastTimeChecked) > timeout {
			fmt.Println(lastTimeChecked)
			go printFile(file)
			file = make([]int, 0)
			lastTimeChecked = time.Now()
		}
	}
}

func printFile(file []int) {
	fmt.Println(file)
}

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	dataBuffer := make(chan message)
	wg := &sync.WaitGroup{}

	wg.Add(2)
	go produceData(ctx, dataBuffer, wg)
	go batchData(ctx, dataBuffer, 5*time.Second, 10, wg)
	wg.Wait()

	fmt.Println("Main done")
}
