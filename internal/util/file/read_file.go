package manage_file

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func ReadFile(filename string) []string {
	var items []string
	file, _ := os.Open(filename)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		items = append(items, strings.ToLower(scanner.Text()))
	}
	return items
}

func ReadFileBytes(path string) []byte {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil
	}
	defer file.Close()

	// Create a new reader for the file
	reader := bufio.NewReader(file)

	// Read the file content into a byte slice
	fileInfo, _ := file.Stat()
	size := fileInfo.Size()
	buffer := make([]byte, size)
	_, err = reader.Read(buffer)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return nil
	}
	return buffer
}
