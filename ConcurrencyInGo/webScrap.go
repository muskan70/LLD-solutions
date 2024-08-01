package main

import (
	"io"
	"log"
	"net/http"
	"sync"
)

type pair struct {
	url     string
	keyword string
	count   int
}

func getKeywordCountFromContent(keyword string, content string) int {
	count := 0
	i := 0
	j := len(keyword)
	for i < len(content) {
		if content[i] == keyword[0] && i+j < len(content) && keyword == content[i:i+j] {
			count++
			i = i + j
		} else {
			i++
		}
	}
	return count
}

func getContentFromUrl(url string) string {
	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		log.Println(err.Error())
		return ""
	}

	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println(err.Error())
		return ""
	}

	// Check server response
	if resp.StatusCode != http.StatusOK {
		log.Println(url + " bad status: " + resp.Status)
		return ""
	}

	return string(bodyBytes)
}

func getKeywordsFromUrl(urls []string, keywords []string, c chan pair, wg *sync.WaitGroup, workerId int) {
	log.Println("workerId:", workerId, "urls:", urls)
	defer wg.Done()
	for i := 0; i < len(urls); i++ {
		content := getContentFromUrl(urls[i])
		count := getKeywordCountFromContent(keywords[i], content)
		c <- pair{url: urls[i], keyword: keywords[i], count: count}
	}

}

func getKeywordOccurences(urls []string, keywords []string, noofWorkers int) []pair {
	var wg *sync.WaitGroup = new(sync.WaitGroup)
	i := 0
	noofUrlsPerWorker := len(urls) / noofWorkers
	workerId := 1
	wg.Add(noofWorkers)
	c := make(chan pair, len(urls))
	for i < len(urls) {
		if workerId == noofWorkers {
			listOfUrls := urls[i:]
			listOfKeywords := keywords[i:]
			go getKeywordsFromUrl(listOfUrls, listOfKeywords, c, wg, workerId)
			break
		}
		listOfUrls := urls[i : i+noofUrlsPerWorker]
		listOfKeywords := keywords[i : i+noofUrlsPerWorker]
		go getKeywordsFromUrl(listOfUrls, listOfKeywords, c, wg, workerId)
		i += noofUrlsPerWorker
		workerId++
	}
	res := []pair{}
	for k := 0; k < len(urls); k++ {
		msg1 := <-c
		res = append(res, msg1)
		log.Println(msg1)
	}
	wg.Wait()
	return res
}
