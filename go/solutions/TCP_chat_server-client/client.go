package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {

	connection, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		panic(err)
	}

	//scanner := bufio.NewScanner(connection)

	var serverMessage string
	var reply string

	inputReader := bufio.NewReader(os.Stdin)
	replyReader := bufio.NewReader(connection)

	for {
		serverMessage, _ = replyReader.ReadString('\n')
		fmt.Println("\n\nServer says: ")
		fmt.Println("\n\n", serverMessage)
		fmt.Println("\n\nReply to server: ")
		reply, _ = inputReader.ReadString('\n')
		fmt.Fprint(connection, reply)
	}

}
