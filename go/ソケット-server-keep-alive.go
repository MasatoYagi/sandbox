package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"net/http/httputil"
	"strings"
	"time"
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
			// Accept後のソケットでコネクションが張られた後に何度もリクエストを受けられるようにループ
			for {
				// 通信がしばらくない(5s)とタイムアウトのエラーでRead()の呼び出しを終了
				conn.SetReadDeadline(time.Now().Add(5 * time.Second))
				// リクエストを読み込む
				request, err := http.ReadRequest(bufio.NewReader(conn))
				if err != nil {
					// タイムアウトもしくはソケットクローズ時は終了
					// それ以外はエラー
					neterr, ok := err.(net.Error) // ダウンキャスト(タイムアウトは、標準のerrorインタフェースの上位互換であるnet.Errorインタフェースの構造体から取得)
					if ok && neterr.Timeout() {
						fmt.Println("Timeout")
						break
					} else if err == io.EOF {
						break
					}
					panic(err)
				}
				// リクエストを表示
				dump, err := httputil.DumpRequest(request, true)
				if err != nil {
					panic(err)
				}
				fmt.Println(string(dump))
				content := "Hello World\n"
				// レスポンスを書き込む
				// HTTP/1.1かつ、COntentLengthの設定が必要
				response := http.Response{
					StatusCode:    200,
					ProtoMajor:    1,
					ProtoMinor:    1,
					ContentLength: int64(len(content)),
					Body:          ioutil.NopCloser(strings.NewReader(content)),
				}
				response.Write(conn)
			}
			conn.Close()
		}()
	}
}
