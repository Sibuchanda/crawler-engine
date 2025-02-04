package modules

import (
	"errors"
	"math/rand/v2"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Range struct {
	Low  uint8
	High uint8
}

type ProbQueue struct {
	Name        string
	Queue       amqp.Queue
	Probability Range
}

type MQ struct {
	channel    *amqp.Channel
	connection *amqp.Connection
	consumer   map[string]<-chan amqp.Delivery
	Queues     map[string]ProbQueue
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

	t.Queues = make(map[string]ProbQueue)

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
func (t *MQ) DeclareQueue(name string, prob Range) (err error) {
	if prob.Low > 100 {
		return errors.New("probability Low value can't be larger than 100 or less than 0")
	}

	if prob.High > 100 {
		return errors.New("probability High value can't be larger than 100 or less than 0")
	}

	queue, err := t.channel.QueueDeclare(
		name,  // Name
		true,  // durable
		false, // Auto-delete
		false, // sharable
		false, // No-wait
		nil,
	)

	new_queue := ProbQueue{
		Name:        name,
		Queue:       queue,
		Probability: prob,
	}

	t.Queues[name] = new_queue
	return
}

// SendMessage publish the message into Queue
func (t *MQ) SendMessage(message []byte, queueName string) (err error) {
	if _, ok := t.Queues[queueName]; !ok {
		return errors.New("queue doesn't exist, Please declare queue first")
	}

	err = t.channel.Publish(
		"",        //exchange
		queueName, //routing key
		false,     //mandatory
		false,     //immediate
		amqp.Publishing{
			ContentType:  "text/plain",
			Body:         message,
			DeliveryMode: amqp.Persistent,
		},
	)

	return
}

// ReceiveMessage receives the message from the Queue
func (t *MQ) ReceiveMessage(queueName string) (message []byte, err error) {
	if _, ok := t.Queues[queueName]; !ok {
		return nil, errors.New("queue doesn't exist, Please declare the queue first")
	}

	var consumer <-chan amqp.Delivery

	if len(t.consumer) == 0 {
		consumer, err = t.channel.Consume(
			queueName, // Queue Name
			"",        // consumer
			false,     // Auto Acknowledgment
			false,     // Exclusive/Sharable
			false,     // No-local
			false,     // No-wait
			nil,       // Arguments
		)

		if err != nil {
			return nil, err
		}

		t.consumer = make(map[string]<-chan amqp.Delivery)
		t.consumer[queueName] = consumer
	} else {
		if ch, ok := t.consumer[queueName]; ok {
			consumer = ch
		} else {
			consumer, err = t.channel.Consume(
				queueName, // Queue Name
				"",        // consumer
				false,     // Auto Acknowledgment
				false,     // Exclusive/Sharable
				false,     // No-local
				false,     // No-wait
				nil,       // Arguments
			)

			if err != nil {
				return nil, err
			}

			t.consumer[queueName] = consumer
		}
	}

	select {
	case msg := <-consumer:
		message = msg.Body
		msg.Ack(true)
		return message, nil
	case <-time.After(3 * time.Second):
		return
	}
}

// PickQueues Function that pick queues according to the priority set into the Queue
func (t *MQ) PickQueues() (queue ProbQueue, err error) {
	randomNumber := uint8(rand.N(101))

	for _, queue := range t.Queues {
		if randomNumber >= queue.Probability.Low && randomNumber <= queue.Probability.High {
			return queue, nil
		}
	}

	return ProbQueue{}, errors.New("probability range not found")
}
