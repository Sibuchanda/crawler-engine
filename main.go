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

	h := mod.Hashing{}
	err = h.Connect("http://127.0.0.1:10000", "v1")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = h.GetNode64(hash)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(h.Result.PORT)
	fmt.Println(h.Result.IP)
}
