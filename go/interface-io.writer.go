package main

import (
	"bytes"
	"fmt"
	"io"
	"net"
	"os"
)

// 低レベルアクセスへの入り口（1）：io.Writer
// https://ascii.jp/elem/000/001/243/1243667/

func main() {
	// # ファイル出力
	// os.File のインスタンスは、os.Create()（新規ファイルの場合）や os.Open()（既存のファイルのオープン）などの関数で作成
	file, err := os.Create("interface-io.writer.txt")
	// Writeメソッドを使う予定のファイルオブジェクトは、書き込み権限がついたos.Create()から作ったものでなくてはなりません。
	// https://zenn.dev/hsaki/books/golang-io-package/viewer/file#%E3%83%95%E3%82%A1%E3%82%A4%E3%83%AB%E3%81%B8%E3%81%AE%E6%9B%B8%E3%81%8D%E8%BE%BC%E3%81%BF(write)
	if err != nil {
		panic(err)
	}
	defer file.Close() // 確実にリソースが確保された後であるエラーハンドリング後に実施すること！！！！
	count, err := file.Write([]byte("ファイルへのWrite\n"))
	if err != nil {
		panic(err)
	}
	fmt.Printf("write %d bytes\n", count)

	// # 標準出力
	os.Stdout.Write([]byte("標準出力へのWrite\n"))

	// # 書かれた内容を記憶しておくバッファ
	// bytes.Buffer: Write() メソッドで書き込まれた内容を淡々とためておいて後でまとめて結果を受け取れる
	var buffer bytes.Buffer
	buffer.Write([]byte("バッファへのWrite-0, "))
	fmt.Println(buffer.String())
	// bytes.Buffer には、特別に文字列を受け取れる WriteString というメソッドもある
	buffer.WriteString("バッファへのWrite-1, ")
	fmt.Println(buffer.String())
	// WriteString は io.Writer のメソッドではないため、他の構造体では使えません。代わりに、次の io.WriteString を使えばキャストは不要
	io.WriteString(&buffer, "バッファへのWrite-2")
	fmt.Println(buffer.String())

	// # インターネットアクセス
	// ## TCP
	// https://zenn.dev/hsaki/books/golang-io-package/viewer/netconn
	conn, _ := net.Dial("tcp", "ascii.jp:80")
	conn.Write([]byte("GET / HTTP/1.1\r\nHost: ascii.jp\r\n\r\n"))
	io.Copy(os.Stdout, conn)

	// ## HTTP(http.ResponseWriter: Webサーバーからブラウザに対してメッセージを"書き込む")
	// func handler(w http.ResponseWriter, r *http.Request) {
	// 	w.Write([]byte("http.ResponseWriter sample"))
	// }
	// http.HandleFunc("/", handler)
	// http.ListenAndServe(":8080", nil)

}
