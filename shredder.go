// Description: A simple file shredder that overwrites a file with random data multiple times before deleting it.
// Author: Piotr Koszur 
// Last updated: 2024-03-24
// Usage: go run shredder.go <file_path>
// Or compile with: go build shredder.go
// Then run: ./shredder <file_path>
package main

import (
	"crypto/rand"
	"fmt"
	"os"
)

var shredCount = 3

func overwriteFileWithRandomData(filePath string, fileSize int64) error {
	data := make([]byte, fileSize)
	rand.Read(data)

	err := os.WriteFile(filePath, data, 0644)
	if err != nil {
		return fmt.Errorf("Failed to write file: %v", err)
	}

	return nil
}

func getFileSize(filePath string) int64 {
	fileInfo, err := os.Stat(filePath)
	if err != nil {
		fmt.Printf("Failed to get file size: %v", err)
	}

	return fileInfo.Size()
}

func shredFile(filePath string) error {
	fmt.Printf("Shredding file: %v\n", filePath)
	fileSize := getFileSize(filePath)

	for i := 0; i < shredCount; i++ {
		overwriteFileWithRandomData(filePath, fileSize)
		fmt.Printf("Iteration %v completed.\n", i+1)
	}

	fmt.Println("Shredding done, deleting file.")

	err := os.Remove(filePath)
	if err != nil {
		return fmt.Errorf("Failed to delete file: %v", err)
	}

	return nil
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide a file path as an argument.")
		os.Exit(1)
	}
	filePath := os.Args[1]

	err := shredFile(filePath)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
}
