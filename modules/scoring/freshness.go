package scoring

import (
	"strings"
	"time"

	"github.com/gocql/gocql"
)

// FreshnessScore determines if a URL has fresh content based on timestamps.
func (t *Cassandra) FreshnessScore(url string, urlLastModified time.Time, previousTime time.Time) bool {

	if urlLastModified.IsZero() && previousTime.IsZero() {
		return false
	}

	if !urlLastModified.IsZero() && previousTime.IsZero() {
		return false
	}

	if urlLastModified.Equal(previousTime) {
		return false
	}

	if urlLastModified.After(previousTime) {
		return true
	}

	return false
}

// CheckDomain checks if the URL belongs to frequently updated sections
func CheckDomain(url string) bool {
	sections := []string{"/news/", "/blog/", "/latest/", "/updates/", "/breaking/"}
	url = strings.ToLower(url)

	for _, section := range sections {
		if strings.Contains(url, section) {
			return true
		}
	}
	return false
}

// UpdateLastModified updates the Last-Modified timestamp for a URL if it's newer.
func (t *Cassandra) UpdateLastModified(url string, newTime time.Time) error {
	if newTime.IsZero() {
		return nil
	}

	var storedTime time.Time
	err := t.session.Query(`SELECT last_modified_time FROM last_modified WHERE url = ?`, url).Scan(&storedTime)

	if err != nil && err != gocql.ErrNotFound {
		return err
	}

	if err == gocql.ErrNotFound || newTime.After(storedTime) {
		return t.session.Query(`INSERT INTO last_modified (url, last_modified_time) VALUES (?, ?)`, url, newTime).Exec()
	}

	return nil
}

// GetLastModified fetches the last modified timestamp of a given URL
func (t *Cassandra) GetLastModified(url string) (time.Time, error) {
	var storedTime time.Time
	err := t.session.Query(`SELECT last_modified_time FROM last_modified WHERE url = ?`, url).Scan(&storedTime)

	if err != nil {
		return time.Time{}, err
	}
	return storedTime, nil
}
