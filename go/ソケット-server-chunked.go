package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/http/httputil"
	"strings"
)

// https://ascii.jp/elem/000/001/276/1276572/
// TCPソケットを使ったHTTPサーバ

// チャンク形式であれば、準備ができた部分からレスポンスを開始できるため、レスポンスの初動が早くなります。
// 特にファイルサイズが大きくなると効果が大きくなります。
// ファイルから細切れに読み込んで少しずつソケットに流していければ、データ全体を保持するために大量のメモリを確保するというオーバーヘッドも減らせます。
// ヘッダーにサイズを入れる必要もないので、最終的なデータのサイズが決まる前に送信を開始することもできます。

// 青空文庫: ごんぎつねより
// http://www.aozora.gr.jp/cards/000121/card628.html
var contents = []string{
	"これは、私わたしが小さいときに、村の茂平もへいというおじいさんからきいたお話です。",
	"むかしは、私たちの村のちかくの、中山なかやまというところに小さなお城があって、",
	"中山さまというおとのさまが、おられたそうです。",
	"その中山から、少しはなれた山の中に、「ごん狐ぎつね」という狐がいました。",
	"ごんは、一人ひとりぼっちの小狐で、しだの一ぱいしげった森の中に穴をほって住んでいました。",
	"そして、夜でも昼でも、あたりの村へ出てきて、いたずらばかりしました。",
}

// クライアントはgzipを受け入れ可能か？
func isGZipAcceptable(request *http.Request) bool {
	return strings.Index(strings.Join(request.Header["Accept-Encoding"], ","), "gzip") != -1
}

// 1セッションの処理をする
func processSession(conn net.Conn) {
	fmt.Printf("Accept %v\n", conn.RemoteAddr())
	defer conn.Close()
	for {
		// リクエストを読み込む
		request, err := http.ReadRequest(bufio.NewReader(conn))
		if err != nil {
			if err == io.EOF {
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
		fmt.Fprintf(conn, strings.Join([]string{
			"HTTP/1.1 200 OK",
			"Content-Type: text/plain",
			"Transfer-Encoding: chunked",
			"", "",
		}, "\r\n"))
		for _, content := range contents {
			bytes := []byte(content)
			fmt.Println(len(bytes))
			fmt.Println(content)
			fmt.Fprintf(conn, "%x\r\n%s\r\n", len(bytes), content)
		}
		fmt.Fprintf(conn, "0\r\n\r\n")
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
