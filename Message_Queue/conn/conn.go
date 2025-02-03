package conn

import (
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

// *---------For Error Handling-----------------
func FailOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s :%s", msg, err)
	}
}

// *----------Connect with RabbitMq-------------
func ConnectRabbitMq() (*amqp.Connection, *amqp.Channel) {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	FailOnError(err, "Failed to connect to RabbitMQ")
	ch, err := conn.Channel()
	FailOnError(err, "Failed to Opened Channel")
	return conn, ch
}
