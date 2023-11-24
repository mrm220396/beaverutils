package main

import (
	"fmt"
	"os"
)

func main() {
	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Printf("Error getting current directory: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(currentDir)
}
