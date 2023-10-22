package main

import (
	"fmt"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func main() {
	csvService := new(ArangoService)
	rpc.Register(csvService)

	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		fmt.Println("Error opening port 1234")

		return
	}
	defer listener.Close()

	fmt.Println("Service RPC CsvService listening at port 1234")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error while accepting the connection :", err)
			continue
		}
		go jsonrpc.ServeConn(conn)
	}
}
