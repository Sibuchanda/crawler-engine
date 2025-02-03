package main

import (
	"crawler-engine/Message_Queue/conn"
	"fmt"
	"math/rand"

	amqp "github.com/rabbitmq/amqp091-go"
)

// *----------Declare a Queue--------------------
func declareQueue(ch *amqp.Channel, queueName string) {
	_, err := ch.QueueDeclare(
		queueName, //name
		true,      //durable
		false,     //Auto-delete
		false,     //sharable
		false,     //No-wait
		nil,
	)
	conn.FailOnError(err, "Failed to Declare queue")
}

// * -----------Publish Message-------------
func publishMessage(ch *amqp.Channel, queueName string, messsage string) {
	err := ch.Publish(
		"",        //exchange
		queueName, //routing key
		false,     //mandatory
		false,     //immediate
		amqp.Publishing{
			ContentType:  "text/plain",
			Body:         []byte(messsage),
			DeliveryMode: amqp.Persistent,
		},
	)
	conn.FailOnError(err, "Failed to Publishing the message")
	fmt.Printf("%s\n", messsage)
}

// *----------Main Function--------------
func main() {
	// Connect To RabbitMq
	conn, ch := conn.ConnectRabbitMq()
	defer conn.Close()
	defer ch.Close()

	declareQueue(ch, "queue1")
	declareQueue(ch, "queue2")
	declareQueue(ch, "queue3")

	url := []string{
		"https://www.stackoverflow.com--1",
		"https://www.google.com--2",
		"https://www.github.com--3",
		"https://www.stackoverflow.com--4",
		"https://www.google.com--5",
		"https://www.github.com--6",
		"https://www.stackoverflow.com--7",
		"https://www.google.com--8",
		"https://www.github.com--9",
	}
	// Choose the queue based on their priority
	for i := 0; i < 9; i++ {
		randomNumber := rand.Float64()

		if randomNumber <= 0.5 {
			publishMessage(ch, "queue1", url[i])
		} else if randomNumber <= 0.8 {
			publishMessage(ch, "queue2", url[i])
		} else {
			publishMessage(ch, "queue3", url[i])
		}
	}

}
