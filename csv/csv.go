package main

import (
	"fmt"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	"strings"
)

type ConverterService struct{}

type Request struct {
	TsvData string
}

type Response struct {
	CsvData string
}

func (c *ConverterService) TsvToCsv(request *Request, response *Response) error {

	csvData := strings.ReplaceAll(request.TsvData, "\t", ",")
	// Reemplazar las nuevas líneas por saltos de línea
	csvData = strings.ReplaceAll(csvData, "\n", "\r\n")

	response.CsvData = csvData
	return nil
}

func main() {
	csvService := new(ConverterService)
	rpc.Register(csvService)

	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		fmt.Println("Error al abrir el puerto 1234:", err)
		return
	}
	defer listener.Close()

	fmt.Println("Servicio RPC CsvService escuchando en el puerto 1234...")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error al aceptar la conexión:", err)
			continue
		}
		go jsonrpc.ServeConn(conn)
	}
}
