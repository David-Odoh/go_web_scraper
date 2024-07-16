package main

import (
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"
)

type FetchResult struct {
	URL  string
	Body string
	Err  error
}

type ProcessedResult struct {
	URL    string
	Length int
	Err    error
}

func Fetcher(urls []string, fetchResults chan<- FetchResult, wg *sync.WaitGroup) {
	defer wg.Done()
	for _, url := range urls {
		resp, err := http.Get(url)
		if err != nil {
			fetchResults <- FetchResult{URL: url, Body: "", Err: err}
			continue
		}
		body, err := io.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			fetchResults <- FetchResult{URL: url, Body: "", Err: err}
			continue
		}
		fetchResults <- FetchResult{URL: url, Body: string(body), Err: nil}
	}
}

func Processor(fetchResults <-chan FetchResult, processedResults chan<- ProcessedResult, wg *sync.WaitGroup) {
	defer wg.Done()
	for fetchResult := range fetchResults {
		if fetchResult.Err != nil {
			processedResults <- ProcessedResult{URL: fetchResult.URL, Length: 0, Err: fetchResult.Err}
			continue
		}
		processedResults <- ProcessedResult{URL: fetchResult.URL, Length: len(fetchResult.Body), Err: nil}
	}
}

func Aggregator(processedResults <-chan ProcessedResult, wg *sync.WaitGroup) {
	defer wg.Done()
	for processedResult := range processedResults {
		if processedResult.Err != nil {
			fmt.Printf("Error processing URL %s: %v\n", processedResult.URL, processedResult.Err)
		} else {
			fmt.Printf("URL: %s, Content Length: %d\n", processedResult.URL, processedResult.Length)
		}
	}
}

func main() {
	urls := []string{
		"http://youtube.com",
		"http://github.com",
		"http://google.com",
	}

	fetchResults := make(chan FetchResult, len(urls))
	processedResults := make(chan ProcessedResult, len(urls))

	var wg sync.WaitGroup

	wg.Add(1)
	go Fetcher(urls, fetchResults, &wg)

	wg.Add(1)
	go Processor(fetchResults, processedResults, &wg)

	wg.Add(1)
	go Aggregator(processedResults, &wg)

	go func() {
		wg.Wait()
		close(fetchResults)
		close(processedResults)
	}()

	time.Sleep(5 * time.Second)
}
