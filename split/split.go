package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	datacloud "github.com/arejula27/dataCloud"
)

const MAX_LINES = 200

type SplitService struct{}

func (c *SplitService) Split(request *datacloud.Request, response *datacloud.SplitResponse) error {

	//Read file in shared storage
	fileContent, err := os.ReadFile(request.Filename)
	if err != nil {
		log.Println("Error while reading file:", err)
		return err

	}
	fileContentStr := string(fileContent)

	//Split the content
	lines := strings.Split(fileContentStr, "\n")

	//check how many files are needed
	numFiles := (len(lines))/MAX_LINES + 1
	for i := 0; i < numFiles; i++ {
		startIndex := i * MAX_LINES
		endIndex := (i + 1) * MAX_LINES
		if endIndex > len(lines) {
			endIndex = len(lines)
		}
		//Create files
		fileName := fmt.Sprintf("%s_%d.csv", request.Filename, i+1)
		response.Files = append(response.Files, fileName)
		file, err := os.Create(fileName)
		if err != nil {
			return err
		}
		defer file.Close()

		// EWrite lines
		for j := startIndex; j < endIndex; j++ {
			_, err := file.WriteString(lines[j] + "\n")
			if err != nil {
				return err
			}
		}
	}

	return nil
}
