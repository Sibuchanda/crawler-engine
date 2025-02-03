package main

import (
	"crawler-engine/Message_Queue/conn"
	"fmt"
	"log"
	"math/rand/v2"

	amqp "github.com/rabbitmq/amqp091-go"
)

// *----------Declare a Queue--------------------
func declareQueue(ch *amqp.Channel, queueName string) {
	_, err := ch.QueueDeclare(
		queueName, //name
		true,      //durable
		false,     //Auto-delete
		false,     //sharable/Exclusive
		false,     //No-wait
		nil,
	)
	conn.FailOnError(err, "Failed to Declare queue")
}

// *------------Consumes Message---------
func consumesMessage(queueName string, ch *amqp.Channel) {

	mssg, err := ch.Consume(
		queueName, //Queue Name
		"",        //consumer
		false,     //Auto Acknowledgment
		false,     //Exclusive/Sharable
		false,     //No-local
		false,     //No-wait
		nil,       //Arguments

	)
	conn.FailOnError(err, "Failed to Consumes Message")

	go func() {
		for d := range mssg {
			log.Printf("Recives message in %s %s", queueName, d.Body)
			d.Ack(false)
			return
		}
	}()

}

// *------------Main Function-----------
func main() {
	// Connect To RabbitMq
	conn, ch := conn.ConnectRabbitMq()
	defer conn.Close()
	defer ch.Close()

	// Declare a queue
	declareQueue(ch, "queue1")
	declareQueue(ch, "queue2")
	declareQueue(ch, "queue3")

	// consumes Message
	randomNumber := rand.Float64()

	if randomNumber <= 0.5 {
		fmt.Println("Choose 50% Priority")
		consumesMessage("queue1", ch)
	} else if randomNumber <= 0.8 {
		fmt.Println("Choose 30% Priority")
		consumesMessage("queue2", ch)
	} else {
		fmt.Println("Choose 20% Priority")
		consumesMessage("queue3", ch)
	}

}
