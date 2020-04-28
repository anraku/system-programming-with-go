package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "example.com:80")
	if err != nil {
		log.Fatal(err)
	}
	conn.Write([]byte("GET / HTTP/1.0\r\nHost: example.com\r\n\r\n"))

	res, err := http.ReadResponse(bufio.NewReader(conn), nil)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	fmt.Println(res.Header)
	io.Copy(os.Stdout, res.Body)
}
