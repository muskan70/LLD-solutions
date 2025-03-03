package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Enter the cron string Input:")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	cronStr := scanner.Text()

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	if err := CronStringParser(cronStr); err != nil {
		fmt.Println(err)
	}
}
