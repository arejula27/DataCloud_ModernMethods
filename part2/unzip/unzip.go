package main

import (
	"os/exec"

	datacloud "github.com/arejula27/dataCloud/part1"
)

type UnzipService struct{}

func (c *UnzipService) Unzip(request *datacloud.Request, response *datacloud.Response) error {

	//Unzip
	cmd := exec.Command("unzip", request.Filename)

	// Ejecutar el comando
	err := cmd.Run()
	if err != nil {

		return err
	}

	return nil
}
