package main

import (
	"bufio"
	"fmt"
	"net"
	"net/http"
	"net/http/httputil"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8888")
	if err != nil {
		panic(err)
	}
	req, err := http.NewRequest("GET", "http://localhost:8888", nil)
	if err != nil {
		panic(err)
	}
	req.Write(conn)
	res, err := http.ReadResponse(bufio.NewReader(conn), req)
	if err != nil {
		panic(err)
	}
	dump, err := httputil.DumpResponse(res, true)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(dump))
}
