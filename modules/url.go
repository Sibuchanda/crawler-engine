package modules

import (
	"io"
	"net/url"

	"golang.org/x/net/html"
)

// CrawlResult holds the results of a single crawl.
type CrawlResult struct {
	URL   string
	Links []string
	Error error
}

// resolveURL resolves a relative URL against a base URL and returns the absolute URL
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

// extractHref Extracts the href attribute from an anchor (<a>) tag
func extractHref(tag html.Token) string {
	for _, attr := range tag.Attr {
		if attr.Key == "href" {
			return attr.Val
		}
	}
	return ""
}

// ExtractURL Extracts all the URL from the input HTML Code
func ExtractURL(body io.Reader, baseURL string) []string {
	tokenizer := html.NewTokenizer(body)
	links := make([]string, 0)

	for token := tokenizer.Next(); token != html.ErrorToken; token = tokenizer.Next() {
		tag := tokenizer.Token()

		if tag.Data != "a" {
			continue
		}

		url := extractHref(tag)
		if url == "" {
			continue
		}

		absoluteURL := resolveURL(url, baseURL)
		if absoluteURL != "" {
			links = append(links, absoluteURL)
		}
	}

	return links
}

// // Crawl fetches the URL and extracts all links from the page.
// func Crawl(targetURL string) CrawlResult {
// 	resp, err := http.Get(targetURL)
// 	if err != nil {
// 		return CrawlResult{URL: targetURL, Links: nil, Error: err}
// 	}
// 	defer resp.Body.Close()

// 	if resp.StatusCode != http.StatusOK {
// 		return CrawlResult{URL: targetURL, Links: nil, Error: fmt.Errorf("HTTP status: %s", resp.Status)}
// 	}

// 	links := extractURL(resp.Body, targetURL)
// 	return CrawlResult{URL: targetURL, Links: links, Error: nil}
// }

// CrawlEngine starts crawling from a seed URL up to a specified depth.
// func CrawlEngine(seedURL string, depth int) {
// 	visited := make(map[string]bool)
// 	queue := []string{seedURL}

// 	for i := 0; i < depth; i++ {
// 		nextQueue := []string{}
// 		for _, url := range queue {
// 			if visited[url] {
// 				continue
// 			}
// 			visited[url] = true

// 			result := Crawl(url)
// 			if result.Error != nil {
// 				// It is a invslid URL we simply ignore this
// 				// fmt.Printf("Error crawling %s: %v\n", url, result.Error)
// 				continue
// 			}

// 			fmt.Printf("Crawled: %s\n", url)
// 			for _, link := range result.Links {
// 				if !visited[link] && strings.HasPrefix(link, "http") {
// 					nextQueue = append(nextQueue, link)
// 				}
// 			}
// 		}
// 		queue = nextQueue
// 	}
// }

// extractLinks parses the HTML document and extracts all valid links.
// func extractLinks(body io.Reader, baseURL string) []string {
// 	var links []string
// 	z := html.NewTokenizer(body)

// 	for {
// 		tt := z.Next()

// 		switch {
// 		case tt == html.ErrorToken:
// 			return links
// 		case tt == html.StartTagToken:
// 			t := z.Token()

// 			if t.Data == "a" {
// 				for _, attr := range t.Attr {
// 					if attr.Key == "href" {
// 						link := resolveURL(attr.Val, baseURL)
// 						if link != "" {
// 							links = append(links, link)
// 						}
// 					}
// 				}
// 			}
// 		}
// 	}
// }
