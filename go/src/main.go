// Package main demonstrates a concurrent recursive web scraper in Go.
// It showcases Go's goroutines and concurrency model.
package main

import (
	"fmt"
	"log"
	"net/http"
	"runtime"
	"strings"
	"time"

	"golang.org/x/net/html"
)

// printSystemInfo prints system details like OS, architecture, CPU count, and current time.
func printSystemInfo() {
	fmt.Println("--- System Info ---")
	fmt.Printf("OS: %s\n", runtime.GOOS)
	fmt.Printf("Architecture: %s\n", runtime.GOARCH)
	fmt.Printf("Number of CPUs: %d\n", runtime.NumCPU())
	fmt.Printf("Current Time: %s\n", time.Now().Format(time.RFC1123))
	fmt.Println("-------------------")
}

// extractLinks extracts all links from a given URL.
func extractLinks(url string) ([]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return nil, err
	}

	var links []string
	var extract func(*html.Node)
	extract = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, attr := range n.Attr {
				if attr.Key == "href" {
					links = append(links, attr.Val)
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			extract(c)
		}
	}
	extract(doc)
	return links, nil
}

// scrapeRecursively scrapes links from a URL recursively up to a specified depth.
// Uses goroutines for concurrent scraping.
func scrapeRecursively(url string, depth int, visited map[string]bool) {
	if depth <= 0 || visited[url] {
		return
	}
	visited[url] = true
	fmt.Printf("Scraping: %s (Depth: %d)\n", url, depth)

	links, err := extractLinks(url)
	if err != nil {
		log.Printf("Error scraping %s: %v\n", url, err)
		return
	}

	for _, link := range links {
		if strings.HasPrefix(link, "http") {
			go scrapeRecursively(link, depth-1, visited)
		}
	}
}

func main() {
	fmt.Println("Hello from Go!")

	printSystemInfo()

	fmt.Println("\nRecursively scraping links from a webpage...")
	visited := make(map[string]bool)
	scrapeRecursively("https://example.com", 2, visited)

	time.Sleep(5 * time.Second) // Wait for goroutines to finish
	fmt.Println("\nExiting...")
}
