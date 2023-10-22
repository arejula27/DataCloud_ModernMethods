package main

import (
	"log"
	"os"
	"strings"

	datacloud "github.com/arejula27/dataCloud/part1"
)

type ConverterService struct{}

func (c *ConverterService) TsvToCsv(request *datacloud.Request, response *datacloud.Response) error {

	//Read file in shared storage
	fileContent, err := os.ReadFile(request.Filename)
	if err != nil {
		log.Println("Error while reading file:", err)
		return err

	}
	fileContentStr := string(fileContent)

	//format the file content
	csvData := strings.ReplaceAll(fileContentStr, "\t", ",")

	newFileName := strings.Replace(request.Filename, ".tsv", ".csv", 1)
	os.WriteFile(newFileName, []byte(csvData), 0644)
	if err != nil {
		log.Println("Error while writing file:", err)
		return err

	}
	return nil
}
