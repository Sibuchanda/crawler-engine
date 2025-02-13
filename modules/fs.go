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
	hasher := xxh3.New()         // xxh3 supports incremental hashing
	buf := make([]byte, 64*1024) // 64KB buffer

	for {
		n, err := data.Read(buf)
		if n > 0 {
			_, _ = hasher.Write(buf[:n]) // Feed chunks into the hasher
		}
		if err == io.EOF {
			break
		}
		if err != nil {
			return 0, err
		}
	}

	return hasher.Sum64(), nil
}

// GetHash128 Get the hash of the io.Reader Content
func GetHash128(data io.Reader) (xxh3.Uint128, error) {
	hasher := xxh3.New()         // xxh3 supports incremental hashing
	buf := make([]byte, 64*1024) // 64KB buffer for efficient processing

	for {
		n, err := data.Read(buf)
		if n > 0 {
			_, _ = hasher.Write(buf[:n]) // Feed chunks into the hasher
		}
		if err == io.EOF {
			break
		}
		if err != nil {
			return xxh3.Uint128{}, err
		}
	}

	return hasher.Sum128(), nil
}
