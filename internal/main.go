package main

import (
	"io"
	"os"
)

func main() {
	filePath := "test.txt"
	newContent := "This is new content.1"

	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	_, err = io.WriteString(file, newContent)
	if err != nil {
		panic(err)
	}

	file.WriteString(newContent + "afsdfqwerqwe")

	io.WriteString(file, newContent+"123123123")
	// os.WriteFile(name string, data []byte, perm os.FileMode)
}
