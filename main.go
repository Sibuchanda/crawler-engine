package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"golang.org/x/net/html"
)

// CrawlResult holds the results of a single crawl.
type CrawlResult struct {
	URL   string
	Links []string
	Error error
}

// Crawl fetches the URL and extracts all links from the page.
func Crawl(targetURL string) CrawlResult {
	resp, err := http.Get(targetURL)
	if err != nil {
		return CrawlResult{URL: targetURL, Error: err}
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return CrawlResult{URL: targetURL, Error: fmt.Errorf("HTTP status: %s", resp.Status)}
	}

	links := extractLinks(resp.Body, targetURL)
	return CrawlResult{URL: targetURL, Links: links}
}

// extractLinks parses the HTML document and extracts all valid links.
func extractLinks(body io.Reader, baseURL string) []string {
	var links []string
	z := html.NewTokenizer(body)

	for {
		tt := z.Next()

		switch {
		case tt == html.ErrorToken:
			return links
		case tt == html.StartTagToken:
			t := z.Token()

			if t.Data == "a" {
				for _, attr := range t.Attr {
					if attr.Key == "href" {
						link := resolveURL(attr.Val, baseURL)
						if link != "" {
							links = append(links, link)
						}
					}
				}
			}
		}
	}
}

// resolveURL resolves a relative URL to an absolute URL based on a base URL.
func resolveURL(href, base string) string {
	parsedBase, err := url.Parse(base)
	if err != nil {
		return ""
	}
	parsedURL, err := url.Parse(href)
	if err != nil {
		return ""
	}
	resolvedURL := parsedBase.ResolveReference(parsedURL)
	return resolvedURL.String()
}

// CrawlEngine starts crawling from a seed URL up to a specified depth.
func CrawlEngine(seedURL string, depth int) {
	visited := make(map[string]bool)
	queue := []string{seedURL}

	for i := 0; i < depth; i++ {
		nextQueue := []string{}
		for _, url := range queue {
			if visited[url] {
				continue
			}
			visited[url] = true

			result := Crawl(url)
			if result.Error != nil {
				// It is a invslid URL we simply ignore this
				// fmt.Printf("Error crawling %s: %v\n", url, result.Error)
				continue
			}

			fmt.Printf("Crawled: %s\n", url)
			for _, link := range result.Links {
				if !visited[link] && strings.HasPrefix(link, "http") {
					nextQueue = append(nextQueue, link)
				}
			}
		}
		queue = nextQueue
	}
}

func main() {
	// Give a valid URL
	seedURL := "https://www.w3schools.com/hTml/html_basic.asp"
	depth := 2

	fmt.Printf("Starting crawl from %s with depth %d\n", seedURL, depth)
	CrawlEngine(seedURL, depth)
}
