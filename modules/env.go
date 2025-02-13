package modules

import (
	"errors"
	"fmt"
	"net/url"
	"os"
	"strconv"
	"strings"
)

// type ServerInfo struct {
// 	IP   string
// 	Port uint16
// 	URI  string
// }

type MinIOCredentials struct {
	Endpoint        string
	AccessKey       string
	SecretAccessKey string
}

type QueueDetails struct {
	IP       string
	Port     uint16
	User     string
	Password string
	URI      string // RabbitMQ AMQP URI
}

type Env struct {
	Queue QueueDetails
	// ConsistentHashing ServerInfo
	MinIO MinIOCredentials
}

// parseAMQPURL extracts details from an AMQP URL
func parseAMQPURL(amqpURL string) (QueueDetails, error) {
	parsedURL, err := url.Parse(amqpURL)
	if err != nil {
		return QueueDetails{}, fmt.Errorf("invalid AMQP URL: %s", amqpURL)
	}

	// Extract user and password
	user := parsedURL.User.Username()
	password, _ := parsedURL.User.Password() // `_` ignores the boolean result

	// Extract host (IP and port)
	hostParts := strings.Split(parsedURL.Host, ":")
	if len(hostParts) != 2 {
		return QueueDetails{}, fmt.Errorf("invalid host format: %s", amqpURL)
	}

	ip := hostParts[0]
	portNum, err := strconv.ParseUint(hostParts[1], 10, 16)
	if err != nil {
		return QueueDetails{}, fmt.Errorf("invalid port number of URL: %s", amqpURL)
	}

	return QueueDetails{
		IP:       ip,
		Port:     uint16(portNum),
		User:     user,
		Password: password,
	}, nil
}

// // parseURL parse url and extract ip and port
// func parseURL(uri string) (ServerInfo, error) {
// 	parsedURL, err := url.Parse(uri)
// 	if err != nil {
// 		return ServerInfo{}, fmt.Errorf("invalid URL: %s", uri)
// 	}

// 	ip := parsedURL.Hostname()
// 	port, err := strconv.ParseUint(parsedURL.Port(), 10, 16)
// 	if err != nil {
// 		return ServerInfo{}, fmt.Errorf("unable to parse port of URL: %s", uri)
// 	}

// 	return ServerInfo{
// 		IP:   ip,
// 		Port: uint16(port),
// 		URI:  uri,
// 	}, nil
// }

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

// // loadConsistentHashing Load the Environment Variable Related to Consistent Hashing
// func (t *Env) loadConsistentHashing() (err error) {
// 	api, exists := os.LookupEnv("CH_API")
// 	if !exists {
// 		return errors.New("CH_API environment doesn't exist")
// 	}

// 	chapi, err := parseURL(api)
// 	if err != nil {
// 		return err
// 	}

// 	t.ConsistentHashing = chapi
// 	return
// }

// loadMinIOEnv Loads the Environment Variable Related to MinIO
func (t *Env) loadMinIOEnv() (err error) {
	endpoint, exists := os.LookupEnv("MINIO_ENDPOINT")
	if !exists {
		return errors.New("MINIO_ENDPOINT environment doesn't exist")
	}

	accessKey, exists := os.LookupEnv("MINIO_ACCESSKEY")
	if !exists {
		return errors.New("MINIO_ACCESSKEY environment doesn't exist")
	}

	secretAccessKey, exists := os.LookupEnv("MINIO_SECRET_ACCESSKEY")
	if !exists {
		return errors.New("MINIO_SECRET_ACCESSKEY environment doesn't exist")
	}

	t.MinIO.Endpoint = endpoint
	t.MinIO.AccessKey = accessKey
	t.MinIO.SecretAccessKey = secretAccessKey
	return
}

// LoadEnv loads values from the Environment and Store into Environment Type
func (t *Env) LoadEnv() (err error) {
	err = t.loadQueues()
	if err != nil {
		return err
	}

	// err = t.loadConsistentHashing()
	// if err != nil {
	// 	return err
	// }

	err = t.loadMinIOEnv()
	if err != nil {
		return err
	}

	return
}
