package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
)

//todo: in here you'll build a webserver that is called up on startup

func main() {
	listen, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	fmt.Println("listening on tcp:8080")
	defer listen.Close()
	for {
		connection, err := listen.Accept()
		if err != nil {
			panic(err)
		}

		io.WriteString(connection, "\n Hello from TCP server\n")
		fmt.Fprintln(connection, "How is your day?")
		fmt.Fprintf(connection, "%v", "Well,I Hope!")

		connection.Close()
	}
}

func handle(connection net.Conn) {
	scanner := bufio.NewScanner(connection)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
	}
	defer connection.Close()
}
