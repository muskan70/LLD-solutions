package main

import "fmt"

func main() {
	urls := []string{
		"https://www.makemytrip.com",
		"https://in.bookmyshow.com/explore/home/faridabad",
		"https://www.paytm.com",
		"https://www.google.com",
		"https://www.razorpay.com",
		"https://www.eventbrite.com",
		"https://www.sumologic.com",
	}
	keywords := []string{
		"makemytrip",
		"bookmyshow",
		"paytm",
		"google",
		"razorpay",
		"eventbrite",
		"sumologic",
	}

	var noofWorkers int
	fmt.Println("Enter No.of Workers")
	fmt.Scanln(&noofWorkers)

	response := getKeywordOccurences(urls, keywords, noofWorkers)
	fmt.Println(response)
}
