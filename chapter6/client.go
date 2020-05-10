package main

import (
	"bufio"
	"fmt"
	"net"
	"net/http"
	"net/http/httputil"
	"sync"
	"time"
)

func main() {
	start := time.Now()
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

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
		}()
	}
	wg.Wait()
	end := time.Now()
	fmt.Printf("%f秒\n", (end.Sub(start)).Seconds())
}
