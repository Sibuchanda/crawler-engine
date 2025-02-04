package main

import (
	mod "crawler-engine/modules"
	"log"
)

func main() {
	// Loading Environment
	isworking := false
	Env := mod.Env{}
	// err := Env.LoadEnv()
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	Env.Queue.IP = "localhost"
	Env.Queue.Port = 5672
	Env.Queue.URI = "amqp://guest:guest@localhost:5672/"
	Env.Queue.User = "guest"
	Env.Queue.Password = "guest"

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
		if len(data) == 0 {
			continue
		}

		url := string(data)

		// Print for once
		if !isworking {
			log.Println("Message received... Processing")
			// isworking = true
		}

		// Downloading HTML Code
		body, err := mod.FetchData(url)
		if err != nil {
			log.Fatalln(err)
		}
		urls := mod.ExtractURL(body, url)
		err = mod.SaveFile(body, "process.html")
		if err != nil {
			log.Fatalln(err)
		}

		log.Printf("Total Got URLs: %d\n", len(urls))

		// runtime.Gosched() // For Preventing 100% CPU Usage
	}
}
