package main

import (
	"log"
	"os"

	datacloud "github.com/arejula27/dataCloud/part1"
)

type TransformService struct{}

func (c *TransformService) Transform(request *datacloud.Request, response *datacloud.Response) error {

	//Read file in shared storage
	_, err := os.ReadFile(request.Filename)
	if err != nil {
		log.Println("Error while reading file:", err)
		return err

	}

	//format the file content
	data := ""

	os.WriteFile(request.Filename, []byte(data), 0644)
	if err != nil {
		log.Println("Error while writing file:", err)
		return err

	}
	return nil
}
