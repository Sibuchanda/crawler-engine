## fs.go Docs

### Connect()

Connect() is a function that helps to connect to the RabbitMQ Server

- Parameters:
  - uri: string (The address of RabbitMQ Server)
- Returns:
  - Returns error (if occurred)

### Disconnect()

Disconnect() is a function that disconnect the RabbitMQ Client and Channel

- Parameters:
- Returns:
  - Returns error (if occurred)

### createChannel()

createChannel() is a function that creates a channel of RabbitMQ Client

- Parameters:
- Returns:
  - Returns error (if occurred)

### DeclareQueue()

DeclareQueue() is a function that helps to declare a Queue, if a Queue doesn't exist then it will be created automatically

- Parameters:
  - name: string (The name of the Queue)
- Returns:
  - Returns error (if occurred)

### SendMessage()

SendMessage() is a function that push message into the Queue

- Parameters:
  - message: []byte (The message which will be stored)
- Returns:
  - Returns error (if occurred)

### ReceiveMessage()

ReceiveMessage() is a function that pop the message from the Queue

- Parameters:
- Returns:
  - Returns []byte containing the message
  - Returns error (if occurred)

### PickQueues()

PickQueues() is a function that Picks a random Queue from the slice, and returns it's Index.

- Parameters:
  - queues: []MQ (Slice of MQ Types containing Queues)
- Returns:
  - Returns random index among the queues, If picking random index failed then it returns -1.
