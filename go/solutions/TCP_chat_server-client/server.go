package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
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

	//scanner := bufio.NewScanner(client)

	var inputString string
	var reply string

	inputReader := bufio.NewReader(os.Stdin)

	replyReader := bufio.NewReader(client)
	for {
		fmt.Println("\n\nPlease enter message to client:\n\n")
		inputString, _ = inputReader.ReadString('\n')
		fmt.Fprint(client, inputString)
		reply, _ = replyReader.ReadString('\n')
		fmt.Println("\n\nClient replies:\n\n", reply)
	}

}
