package modules

import (
	"io"
	"os"

	"github.com/zeebo/xxh3"
)

// SaveFile Saves the io.Reader content into file
func SaveFile(content io.Reader, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, content)
	return err
}

// GetHash Get the hash of the io.Reader content
func GetHash64(data io.Reader) (uint64, error) {
	content, err := io.ReadAll(data)
	if err != nil {
		return 0, err
	}

	return xxh3.Hash(content), nil
}
