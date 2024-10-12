package main

import (
	"fmt"
	"net"
)

// https://ascii.jp/elem/000/001/276/1276572/
// TCPソケットを使ったHTTPクライアント

// ソケットを開いて、サーバが10秒に一回送信するパケットを待って表示する
// クライアントコードは、構成としてはTCPの例におけるサーバと同じ

func main() {
	fmt.Println("Listen tick server at 224.0.0.1:9999")
	address, err := net.ResolveUDPAddr("udp", "224.0.0.1:9999")
	if err != nil {
		panic(err)
	}
	listener, err := net.ListenMulticastUDP("udp", nil, address) // UDPによるマルチキャスト専用の関数
	defer listener.Close()
	buffer := make([]byte, 1500)
	for {
		length, remoteAddress, err := listener.ReadFromUDP(buffer)
		if err != nil {
			panic(err)
		}
		fmt.Printf("Server %v\n", remoteAddress)
		fmt.Printf("Now    %s\n", string(buffer[:length]))
	}
}
