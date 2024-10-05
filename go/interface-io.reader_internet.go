package main

import (
	"io"
	"net"
	"os"
)

// 低レベルアクセスへの入り口（2）：io.Reader前編
// https://ascii.jp/elem/000/001/252/1252961/

func main() {
	// 低級な方法なので使うことはない
	conn, err := net.Dial("tcp", "golang.org:80")
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	_, err = conn.Write([]byte("GET / HTTP/1.1\r\nHost: golang.org\r\nConnection: close\r\n\r\n"))
	if err != nil {
		panic(err)
	}
	io.Copy(os.Stdout, conn) // 生のHTTPの通信内容そのものを読み込む
}
