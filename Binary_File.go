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
	fileInfo, _ := file.Stat()
	fileSize := fileInfo.Size()
	buffer := make([]byte, fileSize)

	_, err = file.Read(buffer)
	if err != nil {
		panic(err)
	}
}
