package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"net/http/httputil"
	"strings"
)

// https://ascii.jp/elem/000/001/276/1276572/
// TCPソケットを使ったHTTPサーバ

func main() {
	listener, err := net.Listen("tcp", "localhost:8888")
	if err != nil {
		panic(err)
	}
	fmt.Println("Server is running at localhost:8888")
	for { // 一度で終了しないためにAccept()を何度も繰り返し呼ぶ
		conn, err := listener.Accept() // net.Connは、io.Reader、io.Writer、io.Closerにタイムアウトを設定するメソッドを追加したインタフェースで、通信のための共通インタフェースとして定義
		if err != nil {
			panic(err)
		}
		// 1リクエスト処理中に他のリクエストのAccept()が行えるように
		// Goroutineを使って非同期にレスポンスを処理する
		go func() { // connを使った読み書き
			fmt.Printf("Accept %v\n", conn.RemoteAddr())
			// リクエストを読み込む
			request, err := http.ReadRequest( // クライアントから送られてきたリクエスト読み込み(HTTPリクエストのヘッダー、メソッド、パスなどの情報を切り出し
				bufio.NewReader(conn))
			if err != nil {
				panic(err)
			}
			dump, err := httputil.DumpRequest(request, true) // io.Readerからバイト列を読み込んで分析してデバッグ出力に出す
			if err != nil {
				panic(err)
			}
			fmt.Println(string(dump))
			// レスポンスを書き込む
			response := http.Response{
				StatusCode: 200,
				ProtoMajor: 1,
				ProtoMinor: 0,
				Body: ioutil.NopCloser(
					strings.NewReader("Hello World\n")),
			}
			response.Write(conn)
			conn.Close()
		}()
	}
}
