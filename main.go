package main

import (
	"fmt"
	"os"
)

func main() {
	data := []byte("echo 'Hello, World!'")
	err := os.WriteFile("hello.sh", data, 0777)
	if err != nil {
		fmt.Println("Error writing file:", err)
		return
	}
	fmt.Println("File written successfully")
}
