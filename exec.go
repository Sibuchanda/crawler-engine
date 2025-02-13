package main

import (
	"bytes"
	mod "crawler-engine/modules"
	"fmt"
	"io"
	"log"
	"strconv"
	"strings"

	"github.com/zeebo/xxh3"
)

// run Starts Execution of the Service
//
// Note: this run function doesn't run in Infinite Loop
func run(mq *mod.MQ, s3 *mod.MinIO, bucketName string) {
	url, err := readURL(mq)
	if err != nil {
		if strings.Compare(err.Error(), "no data") == 0 {
			return
		}
		log.Println(err)
		return
	}
	log.Println(url)

	_, err = extractURLs(url)
	if err != nil {
		if strings.Compare(err.Error(), "unable to save file") == 0 {
			log.Fatalln(err)
		}
		log.Println(err)
		return
	}

	hash := xxh3.HashString128(url)
	urlHash := strconv.FormatUint(hash.Hi, 10) + strconv.FormatUint(hash.Lo, 10) // 128-bit Hash
	err = s3.UploadFile(bucketName, urlHash, "process.html", "text/html")
	if err != nil {
		log.Println("unable to upload file, are you sure upload service is working properly?")
		return
	}
}

// readURL Reading an URL (From Queue)
func readURL(mq *mod.MQ) (resp string, err error) {
	queue, err := mq.PickQueues()
	if err != nil {
		return "", fmt.Errorf("unable to pick random Queues, are you sure everything is setup and working correctly?")
	}

	data, err := mq.ReceiveMessage(queue.Name, 10)
	if err != nil {
		return "", fmt.Errorf("unable to read message from Queues, are you sure queue is accessible?")
	}

	if len(data) == 0 {
		return "", fmt.Errorf("no data")
	}

	return string(data), nil
}

// extractURLs Download and Extract all the URL's from HTML Code
//
// It also stores the HTML Code into Process.html file
func extractURLs(url string) (urls []string, err error) {
	body, err := mod.FetchData(url)
	if err != nil {
		return []string{}, fmt.Errorf("unable to download HTML, are you sure the url is valid?")
	}
	defer body.Close()

	var buf bytes.Buffer
	tee := io.TeeReader(body, &buf)
	err = mod.SaveFile(tee, "process.html")
	if err != nil {
		return []string{}, fmt.Errorf("unable to save file, are you sure the current path have permission to save file?")
	}
	urls = mod.ExtractURL(&buf, url)
	return
}
