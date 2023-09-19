package main

import (
	"log"
	"os/exec"

	datacloud "github.com/arejula27/dataCloud"
)

type UnzipService struct{}

func (c *UnzipService) Unzip(request *datacloud.Request, response *datacloud.Response) error {

	//Unzip
	cmd := exec.Command("unzip", request.Filename)

	// Ejecutar el comando
	err := cmd.Run()
	if err != nil {
		log.Println("Error while running'unzip': %v\n%s", err, output)
		return err
	}

	return nil
}
