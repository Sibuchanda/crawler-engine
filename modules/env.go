package modules

import (
	"errors"
	"net/url"
	"os"
	"strconv"
	"strings"
)

type QueueDetails struct {
	IP       string
	Port     uint16
	User     string
	Password string
	URI      string // RabbitMQ AMQP URI
}

type Env struct {
	Queue QueueDetails
}

// parseAMQPURL extracts details from an AMQP URL
func parseAMQPURL(amqpURL string) (QueueDetails, error) {
	parsedURL, err := url.Parse(amqpURL)
	if err != nil {
		return QueueDetails{}, errors.New("invalid AMQP URL")
	}

	// Extract user and password
	user := parsedURL.User.Username()
	password, _ := parsedURL.User.Password() // `_` ignores the boolean result

	// Extract host (IP and port)
	hostParts := strings.Split(parsedURL.Host, ":")
	if len(hostParts) != 2 {
		return QueueDetails{}, errors.New("invalid host format")
	}

	ip := hostParts[0]
	portNum, err := strconv.ParseUint(hostParts[1], 10, 16)
	if err != nil {
		return QueueDetails{}, errors.New("invalid port number")
	}

	return QueueDetails{
		IP:       ip,
		Port:     uint16(portNum),
		User:     user,
		Password: password,
	}, nil
}

// loadQueues Load the Environment Variable Related to Queues
func (t *Env) loadQueues() (err error) {
	q, exists := os.LookupEnv("QUEUE")
	if !exists {
		return errors.New("QUEUE environment doesn't exist")
	}

	details, err := parseAMQPURL(q)
	if err != nil {
		return err
	}

	t.Queue = details
	t.Queue.URI = q

	return
}

// LoadEnv loads values from the Environment and Store into Environment Type
func (t *Env) LoadEnv() (err error) {
	err = t.loadQueues()
	if err != nil {
		return err
	}

	return
}
