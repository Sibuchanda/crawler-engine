package modules

import (
	"io"
	"os"

	"github.com/zeebo/xxh3"
)

// SaveFile Saves the io.Reader content into file
func SaveFile(content io.Reader, filename string) error {
	body, err := io.ReadAll(content)
	if err != nil {
		return err
	}

	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Write(body)
	if err != nil {
		return err
	}
	return nil
}

// GetHash Get the hash of the io.Reader content
func GetHash64(data io.Reader) (uint64, error) {
	content, err := io.ReadAll(data)
	if err != nil {
		return 0, err
	}

	return xxh3.Hash(content), nil
}
