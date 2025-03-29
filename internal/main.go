package main

import (
	"errors"
	"fmt"
	"os"
)

type fileman interface {
	Read(name string)
	Write(name string, content string)
	Append(name string, content string)
}
type filemanImpl struct {
}

func main() {
	command, name, content := os.Args[1], os.Args[2], os.Args[3]

	filemanImpl := filemanImpl{}

	switch command {
	case "read":
		filemanImpl.Read(name)
	case "write":
		filemanImpl.Write(name, content)
	case "append":
		filemanImpl.Append(name, content)
	case "help":
		fmt.Println("Help")
	default:
		fmt.Println("invalid command, currently supported: read, write, append, help")
	}
}

func (f filemanImpl) Read(name string) {
	data, err := os.ReadFile(name)
	if err != nil {
		fmt.Printf("Error reading file: ", err)
		return
	}

	fileContent := string(data)

	fmt.Println(fileContent)
	return
}

func (f filemanImpl) Write(name string, content string) {
	err := validateContent(content)
	if err != nil {
		fmt.Printf(
			"Invalid content:",
			err,
		)
	}

	err = os.WriteFile(name, []byte(content), 0644)
	if err != nil {
		fmt.Printf("Error writing file: ", err)
		return
	}

	fmt.Print("file written successfully")
	return
}

func (f filemanImpl) Append(name string, content string) {
	err := validateContent(content)
	if err != nil {
		fmt.Printf("Invalid content:", err)
	}

	file, err := os.OpenFile(name, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("Error opening file: ", err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Printf("Error closing file: ", err)
		}
	}(file)

	_, err = file.WriteString(content)
	if err != nil {
		fmt.Printf("Error appending to file: ", err)
	}

	fmt.Printf("file appended successfully")
	return
}

func validateContent(content string) error {
	if len(content) == 0 {
		return errors.New("content is empty")
	}
	return nil
}
