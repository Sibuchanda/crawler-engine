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

### Store the URL's back into Queue

#### Freshness Score Implementation

- Checking for Freshness Keywords in the URL (like /news/, /blog/, /latest/, /updates/, /breaking/).
- Check if page modified recently (If Last Modified header exist)
- If any of the Condition Satisfies then Increase the Score

#### Backlink Metrics

- Check how many Backlinks are there for a specific URL
- Increase the score accordingly (total_backlinks_count \* each_backlink_score)

#### Page Depth

- Count the total number of slashes exist into an URL
- Decrease the score accordingly if number of slashs increases

### Store HTML Code into Persistent Memory

- HTML Input
- Asking for Consistent Hashing where to store the data
  - Declare node IP/Port
  - Input Hash as Parameter (XXH3)
  - Sending GET Request
  - Getting Output of the IP/Port
- Store the data into Persistent memory

### Store the HTML Code Information into Queue

- Take the stored HTML File name
- Store the url, html file name, bucket name into Queue (Where indexing Nodes will access content)

# Environment Variables

| Environment Name         | Required | Description                             |
| ------------------------ | -------- | --------------------------------------- |
| `QUEUE`                  | Yes      | Contains the RabbitMQ AMQP url          |
| `CH_API`                 | Yes      | Contains the Consistent hashing API url |
| `MINIO_ENDPOINT`         | Yes      | Contains MinIO API Endpoint             |
| `MINIO_ACCESSKEY`        | Yes      | Contains MinIO AccessKey                |
| `MINIO_SECRET_ACCESSKEY` | Yes      | Contains MinIO Secret AccessKey         |
