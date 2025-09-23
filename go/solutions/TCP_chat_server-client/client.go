package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {

	connection, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(connection)
	for scanner.Scan() {
		t := scanner.Text()
		fmt.Println("\n\nServer says: ", t)
	}

}
