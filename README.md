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


### Freshness Score Implementation

#### Tasks Implemented in freshness.go

`1`  Checking for Freshness Keywords in the URL

- If the URL belongs to frequently updated sections (like /news/, /blog/, /latest/, /updates/, /breaking/).

- If any of these sections are found, then return true

`2` Handling Last-Modified Header

- The function retrieves the Last-Modified date from the URL.

- It then interacts with Cassandra (via cassandra.go) to manage and update stored timestamps.


#### Conditioons Handling Last-Modified Header

`(i)` First time fetching a URL, no last modified date in Cassandra, and URL also has no modified date → Return nothing.

`(ii)` If Cassandra has no date but URL has modified date → Return +30 and also update  in Cassandra.

`(iii)` If URL has a modified date that is the same as stored in Cassandra → Do not return anything, do not update Cassandra.

`(iv)` If URL has a modified date that is latest than Cassandra’s stored date → Return +30 and update Cassandra.