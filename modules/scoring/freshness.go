package scoring

import (
	"strings"
	"time"
)

// FreshnessScore calculates the freshness score of a URL.
func (t *Cassandra) FreshnessScore(url string, urlLastModified time.Time) (int, error) {
	score := 0

	// Task 1: Check if URL belongs to frequently updated sections
	score += isFrequentlyUpdatedSection(url)

	// Task 2: Handle Last-Modified logic
	storedLastModified, err := t.GetLastModified(url)

	if err != nil { // If no record exists in Cassandra
		// Store the new Last-Modified date if available and return +30
		if !urlLastModified.IsZero() {
			err = t.UpdateLastModified(url, urlLastModified)
			if err != nil {
				return 0, err
			}
			score += 30
		}
	} else if urlLastModified.After(storedLastModified) { // If the URL has a newer date
		// Update the Last-Modified date and return +30
		err = t.UpdateLastModified(url, urlLastModified)
		if err != nil {
			return 0, err
		}
		score += 30
	}

	return score, nil
}

// isFrequentlyUpdatedSection checks if the URL is from sections like /news/, /blog/, etc.
func isFrequentlyUpdatedSection(url string) int {
	sections := []string{"/news/", "/blog/", "/latest/", "/updates/", "/breaking/"}
	url = strings.ToLower(url)

	for _, section := range sections {
		if strings.Contains(url, section) {
			return 30 // Return +30 if the condition satisfies
		}
	}
	return 0 // Return nothing (0) if no condition satisfies
}
