package modules

import (
	"io"
	"os"
)

// Save the io.Reader content into file
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
