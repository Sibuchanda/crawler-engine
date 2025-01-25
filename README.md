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

### Store the URL's back into Queue

### Store HTML Code into Persistent Memory

### Store the HTML Code Information into Queue
