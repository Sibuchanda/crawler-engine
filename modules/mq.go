package modules

import (
	"math/rand/v2"

	amqp "github.com/rabbitmq/amqp091-go"
)

type MQ struct {
	channel    *amqp.Channel
	connection *amqp.Connection
	queue      amqp.Queue
	Priority   [2]int
}

// Connect Connects to RabbitMQ Instance
func (t *MQ) Connect(uri string) (err error) {
	t.connection, err = amqp.Dial(uri)
	if err != nil {
		return err
	}

	err = t.createChannel()
	if err != nil {
		return err
	}

	return
}

// Disconnect Closes the Channel and the Connection
func (t *MQ) Disconnect() (err error) {
	err = t.channel.Close()
	if err != nil {
		return err
	}

	err = t.connection.Close()
	if err != nil {
		return err
	}

	return
}

// createChannel Creates a Channel
func (t *MQ) createChannel() (err error) {
	t.channel, err = t.connection.Channel()
	if err != nil {
		return err
	}
	return
}

// DeclareQueue Declares a Queue
func (t *MQ) DeclareQueue(name string) (err error) {
	t.queue, err = t.channel.QueueDeclare(
		name,  //name
		true,  //durable
		false, //Auto-delete
		false, //sharable
		false, //No-wait
		nil,
	)
	return
}

// SendMessage publish the message into Queue
func (t *MQ) SendMessage(message []byte) (err error) {
	err = t.channel.Publish(
		"",           //exchange
		t.queue.Name, //routing key
		false,        //mandatory
		false,        //immediate
		amqp.Publishing{
			ContentType:  "text/plain",
			Body:         message,
			DeliveryMode: amqp.Persistent,
		},
	)

	return
}

// ReceiveMessage receives the message from the Queue
func (t *MQ) ReceiveMessage() (message []byte, err error) {
	ch, err := t.channel.Consume(
		t.queue.Name, // Queue Name
		"",           // consumer
		true,         // Auto Acknowledgment
		false,        // Exclusive/Sharable
		false,        // No-local
		false,        // No-wait
		nil,          // Arguments
	)

	if err != nil {
		return nil, err
	}
	message = (<-ch).Body

	return
}

// PickQueues Function that pick queues according to the priority set into the Queue
func PickQueues(queues []MQ) int {
	randomNumber := rand.N(101)

	for index, queue := range queues {
		if randomNumber >= queue.Priority[0] && randomNumber <= queue.Priority[1] {
			return index
		}
	}

	return -1
}
