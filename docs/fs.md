## fs.go Docs

### SaveFile()

SaveFile() is a function that takes the [io.Reader](https://pkg.go.dev/io#Reader) object, and store it into the current path with the filename provided.

- Parameters:
  - content: [io.Reader](https://pkg.go.dev/io#Reader) (The Content which will be stored)
  - filename: string (Filename where the content will be stored)
- Returns:
  - Returns error (if occurred)
