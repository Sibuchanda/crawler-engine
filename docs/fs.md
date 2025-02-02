## fs.go Docs

### SaveFile()

SaveFile() is a function that takes the [io.Reader](https://pkg.go.dev/io#Reader) object, and store it into the current path with the filename provided.

- Parameters:
  - content: [io.Reader](https://pkg.go.dev/io#Reader) (The Content which will be stored)
  - filename: string (Filename where the content will be stored)
- Returns:
  - Returns error (if occurred)

### GetHash64()

GetHash() is a function that takes the [io.Reader](https://pkg.go.dev/io#Reader) object, and return the [XXH3](http://www.xxhash.com/) hash of it.

- Parameters:
  - data: [io.Reader](https://pkg.go.dev/io#Reader) (The Content which will be used for hashing)
- Returns:
  - Returns uint64
  - Returns error (if occurred)
