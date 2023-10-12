package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("binaryfile.bin")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	
}
