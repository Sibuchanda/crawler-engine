package main

import (
	mod "crawler-engine/modules"
	"fmt"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("https://example.com/")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("HTTP status: %s\n", resp.Status)
		os.Exit(1)
	}

	result := mod.ExtractURL(resp.Body, "https://example.com/")
	fmt.Println(result)
}
