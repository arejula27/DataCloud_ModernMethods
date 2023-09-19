package main

import (
	"fmt"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

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
