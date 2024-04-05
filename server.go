package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
)

//todo: in here you'll build a webserver that is called up on startup

func main() {
	listen, err := net.Listen("tcp", "localhost:8080")
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
		bs := []byte(ln)
		r13 := rot13(bs)
		fmt.Fprintf(connection, "here is your encrypted msg: %s\n ", r13)
	}
	defer connection.Close()
	fmt.Println("connection to tcp server closed")
}

func rot13(bs []byte) string {
	var r13 = make([]byte, len(bs))
	for i, b := range bs {
		if b <= 109 {
			r13[i] = b + 13
		} else {
			r13[i] = b - 13
		}
	}
	return string(r13)
}

// cookiejar.Jar{} is a thing....didnt know that and want to keep that knowledge for later :)
func client() {
	connection, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		panic(err)
	}
	defer connection.Close()
	reader, err := io.ReadAll(connection)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(reader))
}
