package modules

import (
	"io"
	"net/http"
	"os"
)

func FetchHTML(url string) io.Reader {

	// Send a GET request
	resp, err := http.Get(url)
	if err != nil {
		// fmt.Println("Error fetching the URL:", err)
		return nil
	}
	// defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		// fmt.Println("Error reading the response body:", err)
		return nil
	}

	// Write the HTML content to a file
	file, err := os.Create("process.html")
	if err != nil {
		// fmt.Println("Error creating the file:", err)
		return nil
	}
	defer file.Close()

	_, err = file.Write(body)
	if err != nil {
		// fmt.Println("Error writing to the file:", err)
		return nil
	}
	return resp.Body
	// OR, Print the HTML content
	// fmt.Println(string(body))
}
