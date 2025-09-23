package main

import (
	"fmt"
	"net"
)

func main() {

	listener, err := net.Listen("tcp", "localhost:8000")

	if err != nil {
		panic(err)
	}
	connection, err := listener.Accept()
	connection_handler(connection)
}

func connection_handler(client net.Conn) {
	fmt.Fprint(client, "hej")
}
