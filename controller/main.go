package main

import (
	"fmt"
	"net/rpc/jsonrpc"

	datacloud "github.com/arejula27/dataCloud"
)

func tsvToCsv(filename string) error {
	// Conectar al servidor RPC en el puerto 1234
	client, err := jsonrpc.Dial("tcp", "converter:1234")
	if err != nil {
		fmt.Println("Error while conecting to RPC server:", err)
		return err
	}
	defer client.Close()

	// Call TSV to CSV
	var response datacloud.Response
	request := datacloud.Request{Filename: filename}
	err = client.Call("ConverterService.TsvToCsv", &request, &response)
	if err != nil {
		fmt.Println("Error calling remote method:", err)
		return err
	}
	return nil
}

func main() {

	// Mostrar la respuesta CSV por pantalla

	//Convert tsv to csv
	tsvToCsv("file.tsv")

}
