package main

import (
	"io"
	"os"
)

// 低レベルアクセスへの入り口（1）：io.Writer
// https://ascii.jp/elem/000/001/243/1243667/

// io.Writer を受け取り、書き込まれたデータを加工して別の io.Writer に書き出す
func main() {
	file, err := os.Create("interface-io.writer-multiwriter.txt")
	if err != nil {
		panic(err)
	}
	// ファイルと標準出力の両方に書き込む
	writer := io.MultiWriter(file, os.Stdout)
	io.WriteString(writer, "io.MultiWriter example\n")
}
