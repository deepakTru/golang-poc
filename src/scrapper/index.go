package scrapper

import (
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/PuerkitoBio/goquery"
)

// ScrapeResult holds the result of a scrape operation
type ScrapeResult struct {
	URL   string
	Title string
}

// fetchTitle fetches the title of a webpage
func fetchTitle(url string, wg *sync.WaitGroup, results chan<- ScrapeResult) {
	defer wg.Done()

	// Fetch the webpage
	resp, err := http.Get(url)
	if err != nil {
		log.Printf("Error fetching %s: %v\n", url, err)
		return
	}
	defer resp.Body.Close()

	// Check if the response status code is OK
	if resp.StatusCode != http.StatusOK {
		log.Printf("Error: %s returned status code %d\n", url, resp.StatusCode)
		return
	}

	// Parse the HTML using goquery
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Printf("Error parsing HTML for %s: %v\n", url, err)
		return
	}

	// Extract the title
	title := doc.Find("title").Text()

	// Send the result to the channel
	results <- ScrapeResult{URL: url, Title: title}
}

func WebScrapper() {
	// List of URLs to scrape
	urls := []string{
		"https://www.google.com",
		"https://www.github.com",
		"https://www.stackoverflow.com",
		"https://www.reddit.com",
		"https://www.wikipedia.org",
	}

	// Create a WaitGroup to wait for all goroutines to finish
	var wg sync.WaitGroup

	// Create a channel to collect results
	results := make(chan ScrapeResult, len(urls))

	// Launch a goroutine for each URL
	for _, url := range urls {
		wg.Add(1)
		go fetchTitle(url, &wg, results)
	}

	// Wait for all goroutines to finish
	wg.Wait()

	// Close the results channel
	close(results)

	// Print the results
	fmt.Println("Scraping Results:")
	for result := range results {
		fmt.Printf("URL: %s\nTitle: %s\n\n", result.URL, result.Title)
	}
}
