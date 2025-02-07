package main

import (
	"bytes"
	mod "crawler-engine/modules"
	"io"
	"log"
	"strconv"

	"github.com/zeebo/xxh3"
)

func main() {
	// Loading Environment
	isworking := false
	Env := mod.Env{}
	err := Env.LoadEnv()
	if err != nil {
		log.Fatalln(err)
	}

	// Setting Up Queues
	var mqserver mod.MQ
	mqserver.Connect(Env.Queue.URI)
	defer mqserver.Disconnect()

	mqserver.DeclareQueue("queue1", mod.Range{Low: 0, High: 50})
	mqserver.DeclareQueue("queue2", mod.Range{Low: 51, High: 80})
	mqserver.DeclareQueue("queue3", mod.Range{Low: 81, High: 100})

	// Main Process
	log.Println("Service is started... Waiting for Messages")
	for {
		// Reading an URL (From Queue)
		queue, err := mqserver.PickQueues()
		if err != nil {
			log.Fatalln(err)
		}

		data, err := mqserver.ReceiveMessage(queue.Name)
		if err != nil {
			log.Fatalln(err)
		}

		// Skip to next iteration if the current queue is empty
		if len(data) == 0 {
			continue
		}

		url := string(data)

		// Print for once
		if !isworking {
			log.Println("Message received... Processing")
			isworking = true
		}

		// Downloading HTML Code
		body, err := mod.FetchData(url)
		if err != nil {
			log.Fatalln(err)
		}

		var buf bytes.Buffer
		tee := io.TeeReader(body, &buf)
		err = mod.SaveFile(tee, "process.html")
		if err != nil {
			log.Fatalln(err)
		}
		urls := mod.ExtractURL(&buf, url)
		body.Close()

		log.Printf("Total Got URLs: %d\n", len(urls))

		// Storing HTML file into Persistent Memory
		s3 := mod.MinIO{}
		chash := mod.Hashing{}

		err = chash.Connect(Env.ConsistentHashing.URI, "v1")
		if err != nil {
			log.Fatalln(err)
		}

		urlhash := xxh3.HashString(url)

		err = chash.GetNode64(urlhash)
		if err != nil {
			log.Fatalln(err)
		}

		err = s3.Connect(Env.MinIO.Endpoint, Env.MinIO.AccessKey, Env.MinIO.SecretAccessKey)
		if err != nil {
			log.Fatalln(err)
		}

		err = s3.UploadFile("storage", strconv.FormatUint(urlhash, 10), "process.html", "text/html")
		if err != nil {
			log.Fatalln(err)
		}

		// Storing Information into Queue for Indexing Node to Process

		// runtime.Gosched() // For Preventing 100% CPU Usage
	}
}
