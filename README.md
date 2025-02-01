![Building](https://img.shields.io/github/actions/workflow/status/BiltuDas1/crawler-engine/go-build.yml?label=Build&logo=textpattern&style=flat-square&logoColor=white)
![Testing](https://img.shields.io/github/actions/workflow/status/BiltuDas1/crawler-engine/go-test.yml?label=Test&logo=speedtest&style=flat-square&logoColor=white)

# Crawler Engine

A simple Web Crawler for a model Search Engines like Google

# Algorithm

### Setting Up Queues (Reading an URL)

### Downloading HTML Code

- Take Input URL
- Download the HTML Code of the URL and save it with the name `process.html`
- Return the [io.Reader](https://pkg.go.dev/io#Reader) Object file of the downloaded HTML

### Extracting URL's from the HTML Code

- Input HTML Code
- Extract Anchor Tag
- Extract href attribute value from the Anchor Tag
- Return the href attribute value

### Store HTML Code into Persistent Memory

- HTML Input
- Asking for Consistent Hashing where to store the data
  - Declare node IP/Port
  - Input Hash as Parameter (XXH3)
  - Sending GET Request
  - Getting Output of the IP/Port
- Store the data into Persistent memory

### Store the URL's back into Queue

### Store HTML Code into Persistent Memory

### Store the HTML Code Information into Queue
