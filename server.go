package main

import (
	"bufio"
	"fmt"
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

		go handle(connection)
	}
}

func handle(connection net.Conn) {
	scanner := bufio.NewScanner(connection)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		fmt.Fprintf(connection, "I heard you say: %s \n", ln)
	}
	defer connection.Close()
	fmt.Println("connection to tcp server closed")
}
