package main

import (
	"log"
	"os"
)

func main() {

	file_path := os.Args[1]
	_, err := os.ReadFile(file_path)
	if err != nil {
		log.Fatal("Error while reading file:", err)

	}

	//format the file content
	data := ""
	filename := ""

	os.WriteFile(filename, []byte(data), 0644)
	if err != nil {
		log.Fatal("Error while writing file:", err)

	}
	os.WriteFile("/p1.txt", []byte(filename), 0644)
	if err != nil {
		log.Fatal("Error while writing file:", err)

	}

}
