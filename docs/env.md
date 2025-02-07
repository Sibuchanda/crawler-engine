## env.go Docs

### parseAMQPURL()

parseAMQPURL() is a function that takes the AMQP URL and Returns QueueDetails object of it

- Parameters:
  - amqpURL: string (The AMQP URL)
- Returns:
  - Returns QueueDetails object
  - Returns error (if occurred)

### parseURL()

parseURL() is a function that takes input URL, and Returns ServerInfo Object of it

- Parameters:
  - uri: string
- Returns:
  - Returns ServerInfo object
  - Returns error (if occurred)

### loadQueues()

loadQueues() is a method that loads the Queue details from the Environment

- Parameters:
- Returns:
  - Returns error (if occurred)

### loadConsistentHashing()

loadConsistentHashing() is a method that loads the Consistent Hashing details from the Environment

- Parameters:
- Returns:
  - Returns error (if occurred)

### loadMinIOEnv()

loadMinIOEnv() is a method that loads the MinIO details from the Environment

- Parameters:
- Returns:
  - Returns error (if occurred)

### LoadEnv()

LoadEnv() is a public method that loads all the required details from the Environment into Env object

- Parameters:
- Returns:
  - Returns error (if occurred)
