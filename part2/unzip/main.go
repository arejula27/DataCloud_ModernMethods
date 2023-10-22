package main

import (
	"log"
	"os"
	"os/exec"
)

func main() {

	file_path := os.Args[1]

	//Unzip
	cmd := exec.Command("unzip", file_path)
	// Ejecutar el comando
	err := cmd.Run()
	if err != nil {

		log.Fatalln(err)
	}
	filename := ""
	if err != nil {
		log.Fatal("Error while writing file:", err)

	}
	os.WriteFile("/p1.txt", []byte(filename), 0644)
	if err != nil {
		log.Fatal("Error while writing file:", err)

	}

}
