## mq.go Docs

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
  - prob: Range (Probability of Choosing Queue)
- Returns:
  - Returns error (if occurred)

### SendMessage()

SendMessage() is a function that push message into the Queue

- Parameters:
  - message: []byte (The message which will be stored)
  - queueName: string (The name of the queue)
- Returns:
  - Returns error (if occurred)

### ReceiveMessage()

ReceiveMessage() is a function that pop the message from the Queue

- Parameters:
  - queueName: string (The name of the queue)
- Returns:
  - Returns []byte containing the message
  - Returns error (if occurred)

### PickQueues()

PickQueues() is a function that Picks a random Queue from the queues which is declared using DeclareQueue()

- Parameters:
- Returns:
  - Returns the ProbQueue Object
  - Returns error (if occurred)
