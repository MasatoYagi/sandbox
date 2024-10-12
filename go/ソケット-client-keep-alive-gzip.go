package main

import (
	"bufio"
	"compress/gzip"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"strings"
)

// https://ascii.jp/elem/000/001/276/1276572/
// TCPソケットを使ったHTTPクライアント

func main() {
	sendMessages := []string{
		"ASCII",
		"PROGRAMMING",
		"PLUS",
	}
	current := 0
	var conn net.Conn = nil
	// リトライ用にループで全体を囲う
	for {
		var err error
		// まだコネクションを張ってない / エラーでリトライ時はDialから行う
		if conn == nil {
			conn, err = net.Dial("tcp", "localhost:8888")
			if err != nil {
				panic(err)
			}
			fmt.Printf("Access: %d\n", current)
		}
		// POSTで文字列を送るリクエストを作成
		request, err := http.NewRequest(
			"POST",
			"http://localhost:8888",
			strings.NewReader(sendMessages[current]),
		)
		if err != nil {
			panic(err)
		}
		request.Header.Set("Accept-Encoding", "gzip") // 自分が対応しているアルゴリズムを宣言
		err = request.Write(conn)
		if err != nil {
			panic(err)
		}
		// サーバから読み込む。タイムアウトはここでエラーになるのでリトライ
		response, err := http.ReadResponse(bufio.NewReader(conn), request)
		if err != nil {
			fmt.Println("Retry")
			conn = nil
			continue
		}
		// 結果を表示
		dump, err := httputil.DumpResponse(response, false) // 圧縮された内容を理解してくれないため、2番目のパラメータにfalseを設定してボディを無視
		if err != nil {
			panic(err)
		}
		fmt.Println(string(dump))
		defer response.Body.Close()
		if response.Header.Get("Content-Encoding") == "gzip" { // Accept-Encodingで表明した圧縮メソッドにサーバが対応していたかどうかは、Content-Encodingヘッダーを見ればわかる。表明したアルゴリズムに対応していれば、そのアルゴリズム名がそのまま返ってくる。 今回は1種類だけだが、複数の候補を提示してサーバに選ばせることもできる。 実際、HTTPには、1応答の間に最適なアルゴリズムのネゴシエーションを行う仕組みが備わっている。
			reader, err := gzip.NewReader(response.Body)
			if err != nil {
				panic(err)
			}
			io.Copy(os.Stdout, reader)
			reader.Close()
		} else {
			io.Copy(os.Stdout, response.Body)
		}
		// 全部送信完了していれば終了
		current++
		if current == len(sendMessages) {
			break
		}
	}
}
