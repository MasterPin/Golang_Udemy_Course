package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"net/http"
	"strings"
)

//todo: in here you'll build a webserver that is called up on startup

func main() {

	listen, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		panic(err)
	}
	fmt.Println("listening on tcp:8080")
	defer listen.Close()

	mux := http.NewServeMux()
	mux.HandleFunc("/query", query)
	http.ListenAndServe("localhost:8080", mux)
	for {
		connection, err := listen.Accept()
		if err != nil {
			panic(err)
		}

		go handle(connection)
	}

}

func query(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "you reached the query!")
	r.Body.Read([]byte("you reached the query!"))
}

func handle(connection net.Conn) {
	defer connection.Close()

	request(connection)

	respond(connection)

}

func request(conn net.Conn) {
	i := 0
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		if i == 0 {
			//	request line
			httpMethod := strings.Fields(ln)[0]
			uri := strings.Fields(ln)[1]
			fmt.Println("***METHOD", httpMethod)
			fmt.Println("***URI", uri)
		}
		if ln == "" {
			//	header are done
			break
		}
		i++
	}
}

func respond(conn net.Conn) {
	body := "this is a text"
	fmt.Fprintf(conn, body)
	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprint(conn, "Content-Type: test/html\r\n")
	fmt.Fprint(conn, "\r\n")
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
