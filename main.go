package main

import (
	mod "crawler-engine/modules"
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("test.html")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var hash uint64
	hash, err = mod.GetHash64(f)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(hash)
}
