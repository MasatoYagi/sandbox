package main

import (
	"fmt"
	"net"
)

// https://ascii.jp/elem/000/001/411/1411547/
// UDPソケットを使ったHTTPサーバ

// TCPの通信例と比べるとステップ数が減っていて、そのぶんだけシンプルになっている
func main() {
	fmt.Println("Server is running at localhost:8888")
	// net.ListenPacket()を呼ぶと、net.Listen()のような「クライアントを待つ」インタフェースではなく、
	// データ送受信のためのnet.PacketConnというインタフェースが即座に返されます。
	// このnet.PacketConnオブジェクトもio.Readerインタフェースを実装しているため、
	// 圧縮やファイル入出力などの高度なAPIと簡単に接続できる
	conn, err := net.ListenPacket("udp", "localshot:8888")
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	buffer := make([]byte, 1500)
	for {
		// net.PacketConnは、サーバ側でクライアントを知らない状態で開かれるソケットなので、 このインタフェースを使ってサーバから先にメッセージを送ることはできない
		// サーバには、クライアントからリクエストがあったときに、初めてクライアントのアドレスがわかる
		// ReadFrom()では、TCPのときに紹介した「データの終了を探りながら受信」といった高度な読み込みはできない。
		// そのため、データサイズが決まらないデータに対しては、フレームサイズ分のバッファや、期待されるデータの最大サイズ分のバッファを作り、
		// そこにデータをまとめて読み込むことになる。
		// あるいは、バイナリ形式のデータにしてヘッダにデータ長などを格納しておき、そこまで先読みしてから必要なバッファを確保して読み込む、といったコードになる
		length, remoteAddress, err := conn.ReadFrom(buffer)
		if err != nil {
			panic(err)
		}
		fmt.Printf("Received from %v: %v\n", remoteAddress, string(buffer[:length]))
		_, err = conn.WriteTo([]byte("Hello from Server"), remoteAddress) // ReadFrom()で取得したアドレスに対しては、net.PacketConnインタフェースのWriteTo()メソッドを使ってデータを返送できる
		if err != nil {
			panic(err)
		}
	}
}
