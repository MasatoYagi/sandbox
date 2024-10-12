package main

import (
	"bufio"
	"bytes"
	"compress/gzip"
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

// ヘッダーは圧縮されません。 そのため、少量のデータを通信するほど効率が悪くなります。
// 20バイト足らずのサンプルの文字列ではgzipのオーバーヘッドの方が大きく、サイズが倍増してしまっていますが、
// 大きいサイズになれば効果が出てきます。 ヘッダーの圧縮はHTTP/2になって初めて導入されました。
// なお、HTTPで圧縮されるのはレスポンスのボディだけで、リクエストのボディの圧縮はありません。

// クライアントはgzipを受け入れ可能か？
func isGZipAcceptable(request *http.Request) bool {
	return strings.Index(strings.Join(request.Header["Accept-Encoding"], ","), "gzip") != -1
}

// 1セッションの処理をする
func processSession(conn net.Conn) {
	fmt.Printf("Accept %v\n", conn.RemoteAddr())
	defer conn.Close()
	for {
		conn.SetReadDeadline(time.Now().Add(5 * time.Second))
		// リクエストを読み込む
		request, err := http.ReadRequest(bufio.NewReader(conn))
		if err != nil {
			neterr, ok := err.(net.Error) // ダウンキャスト
			if ok && neterr.Timeout() {
				fmt.Println("Timeout")
				break
			} else if err == io.EOF {
				break
			}
			panic(err)
		}
		dump, err := httputil.DumpRequest(request, true)
		if err != nil {
			panic(err)
		}
		fmt.Println(string(dump))
		// レスポンスを書き込む
		response := http.Response{
			StatusCode: 200,
			ProtoMajor: 1,
			ProtoMinor: 1,
			Header:     make(http.Header),
		}
		if isGZipAcceptable(request) {
			content := "Hello World (gzipped)\n"
			// コンテンツをgzip化して転送
			var buffer bytes.Buffer
			writer := gzip.NewWriter(&buffer)
			io.WriteString(writer, content)
			writer.Close()
			response.Body = ioutil.NopCloser(&buffer)
			response.ContentLength = int64(buffer.Len())
			response.Header.Set("Content-Encoding", "gzip")
		} else {
			content := "Hello World\n"
			response.Body = ioutil.NopCloser(strings.NewReader(content))
			response.ContentLength = int64(len(content))
		}
		response.Write(conn)
	}
}

func main() {
	listener, err := net.Listen("tcp", "localhost:8888")
	if err != nil {
		panic(err)
	}
	fmt.Println("Server is running at localhost:8888")
	for {
		conn, err := listener.Accept()
		if err != nil {
			panic(err)
		}
		go processSession(conn)
	}
}

// func main() {
// 	listener, err := net.Listen("tcp", "localhost:8888")
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Println("Server is running at localhost:8888")
// 	for { // 一度で終了しないためにAccept()を何度も繰り返し呼ぶ
// 		conn, err := listener.Accept() // net.Connは、io.Reader、io.Writer、io.Closerにタイムアウトを設定するメソッドを追加したインタフェースで、通信のための共通インタフェースとして定義
// 		if err != nil {
// 			panic(err)
// 		}
// 		// 1リクエスト処理中に他のリクエストのAccept()が行えるように
// 		// Goroutineを使って非同期にレスポンスを処理する
// 		go func() { // connを使った読み書き
// 			fmt.Printf("Accept %v\n", conn.RemoteAddr())
// 			// Accept後のソケットでコネクションが張られた後に何度もリクエストを受けられるようにループ
// 			for {
// 				// 通信がしばらくない(5s)とタイムアウトのエラーでRead()の呼び出しを終了
// 				conn.SetReadDeadline(time.Now().Add(5 * time.Second))
// 				// リクエストを読み込む
// 				request, err := http.ReadRequest(bufio.NewReader(conn))
// 				if err != nil {
// 					// タイムアウトもしくはソケットクローズ時は終了
// 					// それ以外はエラー
// 					neterr, ok := err.(net.Error) // ダウンキャスト(タイムアウトは、標準のerrorインタフェースの上位互換であるnet.Errorインタフェースの構造体から取得)
// 					if ok && neterr.Timeout() {
// 						fmt.Println("Timeout")
// 						break
// 					} else if err == io.EOF {
// 						break
// 					}
// 					panic(err)
// 				}
// 				// リクエストを表示
// 				dump, err := httputil.DumpRequest(request, true)
// 				if err != nil {
// 					panic(err)
// 				}
// 				fmt.Println(string(dump))
// 				content := "Hello World\n"
// 				// レスポンスを書き込む
// 				// HTTP/1.1かつ、COntentLengthの設定が必要
// 				response := http.Response{
// 					StatusCode:    200,
// 					ProtoMajor:    1,
// 					ProtoMinor:    1,
// 					ContentLength: int64(len(content)),
// 					Body:          ioutil.NopCloser(strings.NewReader(content)),
// 				}
// 				response.Write(conn)
// 			}
// 			conn.Close()
// 		}()
// 	}
// }
