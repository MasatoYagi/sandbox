package main

import (
	"fmt"
	"net"
)

// https://ascii.jp/elem/000/001/276/1276572/
// TCPソケットを使ったHTTPクライアント

// クライアントでは相手がわかった上でDial()するので、TCPの場合と同じようにio.Reader、io.Writerインタフェースのまま使うこともできる。

func main() {
	conn, err := net.Dial("udp4", "localhost:8888")
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	fmt.Println("Sending to server")
	_, err = conn.Write([]byte("Hello from Client"))
	if err != nil {
		panic(err)
	}
	fmt.Println("Receiving from server")
	buffer := make([]byte, 1500)
	length, err := conn.Read(buffer)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Received: %s\n", string(buffer[:length]))
}
