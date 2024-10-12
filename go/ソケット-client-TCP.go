package main

import (
	"bufio"
	"fmt"
	"net"
	"net/http"
	"net/http/httputil"
)

// https://ascii.jp/elem/000/001/276/1276572/
// TCPソケットを使ったHTTPクライアント

func main() {
	conn, err := net.Dial("tcp", "localhost:8888")
	if err != nil {
		panic(err)
	}
	request, err := http.NewRequest(
		"GET", "http://localhost:8888", nil)
	if err != nil {
		panic(err)
	}
	request.Write(conn)
	// http.Response構造体はWrite()メソッドを持っているので、レスポンスのコンテンツをio.Writerに直接書き込める
	response, err := http.ReadResponse(bufio.NewReader(conn), request)
	if err != nil {
		panic(err)
	}
	dump, err := httputil.DumpResponse(response, true)
	if err != nil {
		panic(err)
	}
	fmt.Println(dump)
	fmt.Println(string(dump))
}
