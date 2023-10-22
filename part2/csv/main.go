package main

import (
	"log"
	"os"
	"strings"
)

func main() {

	file_path := os.Args[1]
	fileContent, err := os.ReadFile(file_path)
	if err != nil {
		log.Fatal("Error while reading file:", err)

	}

	//format the file content
	fileContentStr := string(fileContent)

	//format the file content
	csvData := strings.ReplaceAll(fileContentStr, "\t", ",")
	filename := ""
	os.WriteFile(filename, []byte(csvData), 0644)
	if err != nil {
		log.Fatal("Error while writing file:", err)

	}
	os.WriteFile("/p1.txt", []byte(filename), 0644)
	if err != nil {
		log.Fatal("Error while writing file:", err)

	}

}
