package main

import (
	"fmt"
	"io"
	"os"
)

type fileWriter struct{}

func main() {
	filePath := os.Args[1]
	f, err := os.Open(filePath)

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	fw := fileWriter{}

	io.Copy(fw, f)
}

func (fileWriter) Write(b []byte) (n int, err error) {
	fmt.Println(string(b))
	return len(b), nil
}
