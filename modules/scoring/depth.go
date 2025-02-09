package scoring

import (
	"net/url"
	"strings"
)

// CountSlashes Counts the depth of the URL
func CountSlashes(uri string) (counts uint8, err error) {
	baseURL, err := url.Parse(uri)
	if err != nil {
		return 0, err
	}
	counts = uint8(strings.Count(baseURL.Path, "/"))

	if strings.HasSuffix(baseURL.Path, "/") {
		counts--
	}

	return
}
