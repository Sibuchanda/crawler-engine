package main

import (
	mod "crawler-engine/modules"
	"log"
	"runtime"
)

func main() {
	// Loading Environment
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

	// Setting Up Persistent Storage
	s3 := mod.MinIO{}
	err = s3.Connect(Env.MinIO.Endpoint, Env.MinIO.AccessKey, Env.MinIO.SecretAccessKey)
	if err != nil {
		log.Fatalln(err)
	}

	// Main Process
	log.Println("Service is started...")
	for {
		run(&mqserver, &s3, "storage")
		runtime.Gosched() // For Preventing 100% CPU Usage
	}
}
