package main

import (
	"fmt"
	"net/rpc/jsonrpc"

	datacloud "github.com/arejula27/dataCloud"
)

func unzip(filename string) error {

	// Conect to  RPC server  at port 1234
	client, err := jsonrpc.Dial("tcp", "unzip:1234")
	if err != nil {
		fmt.Println("Error while conecting to RPC server:", err)
		return err
	}
	defer client.Close()

	// Call Unzip
	var response datacloud.Response
	request := datacloud.Request{Filename: filename}
	err = client.Call("UnzipService.Unzip", &request, &response)
	if err != nil {
		fmt.Println("Error calling remote method:", err)
		return err
	}
	return nil

}

func tsvToCsv(filename string) error {
	// Conect to  RPC server  at port 1234
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
func split(filename string) ([]string, error) {
	// Conect to  RPC server  at port 1234
	client, err := jsonrpc.Dial("tcp", "converter:1234")
	if err != nil {
		fmt.Println("Error while conecting to RPC server:", err)
		return nil, err
	}
	defer client.Close()

	// Call Split
	response := datacloud.SplitResponse{Files: []string{}}
	request := datacloud.Request{Filename: filename}
	err = client.Call("SplitService.Split", &request, &response)
	if err != nil {
		fmt.Println("Error calling remote method:", err)
		return []string{}, err
	}
	return response.Files, nil

}

func main() {

	// Extact file
	unzip("file.zip")
	//Convert tsv to csv
	tsvToCsv("file.tsv")

	//Split
	files, _ := split("file.csv")

}
