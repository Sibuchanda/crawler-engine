## url.go Docs

### resolveURL()

resolveURL() is a function that takes the base URL and path or other full url, and give absolute URL as output  
It will always give absolute URL as output in every condition. Otherwise it will return empty string.

- Parameters:
  - ref: string (Take URL Path or Full URL as Input)
  - base: string (Take the base URL as Input)
- Returns:
  - Returns Absolute URL

### extractHref()

extractHref() is a function that extract the href attribute from the anchor tag, and return the value.  
It doesn't check if the input tag is the anchor tag or not. Returns Empty String if no href attribute found.

- Parameters:
  - tag: [html.Token](https://pkg.go.dev/golang.org/x/net@v0.33.0/html#Token) (html tag)
- Returns:
  - Returns string value

### ExtractURL()

ExtractURL() is a public function that extracts all the URL from the html body which is passed to the function. It doesn't handle any kind of error, or doesn't do any kind of validation if the input data is HTML or not.

- Parameters:
  - body: [io.Reader](https://pkg.go.dev/io#Reader) (Whole HTML Code)
  - baseURL: string (The original URL, which HTML body is taken as input)
- Returns:
  - Returns slice of string, contains all the extracted URLs

### FetchData()

FetchData() is a public function that takes the input URL, and download the HTML content from the URL and returns a [io.Reader](https://pkg.go.dev/io#Reader) Object of it.

- Parameters:
  - url: string (The URL Which content will be accessed)
- Returns:
  - Returns [io.Reader](https://pkg.go.dev/io#Reader) which contains the HTML Data
  - Returns error if any error occurred
