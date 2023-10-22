package main

import (
	"log"
	"os"
	"strings"
)

func split(path string) []string {
	//TODO: Implement this function
	//wirte file
	return []string{}

}

func main() {

	file_path := os.Args[1]

	files_path_slice := split(file_path)
	//Join the slice files_path into a string with the separator "\n"
	files_path := strings.Join(files_path_slice, "\n")

	err := os.WriteFile("/p1.txt", []byte(files_path), 0644)
	if err != nil {
		log.Fatal("Error while writing file:", err)

	}

}
