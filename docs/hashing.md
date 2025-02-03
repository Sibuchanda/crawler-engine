## hashing.go Docs

### Connect()

Connect() is a public method that helps to connect with the Consistent Hashing Server

- Parameters:
  - address: string (The URL where the Consistent Hashing Server is exposed)
  - apiVersion: string (The API Version of the Consistent Hashing, Supported v1 only)
- Returns:
  - Returns error (if occurred)

### parseData()

parseData() is a method that take the byte JSON data input and extract the IP and Port from it and Store into NodeDetails

- Parameters:
  - data: []byte (byte slice of json response)
- Returns:
  - Returns error (if occurred)

### GetNode64()

GetNode64() is a public method that fetches the NodeDetails where to store the content depending on the hash

- Parameters:
  - hash: uint64 ([XXH3](https://xxhash.com/) Hash of the content which will be stored)
- Returns:
  - Returns error (if occurred)
