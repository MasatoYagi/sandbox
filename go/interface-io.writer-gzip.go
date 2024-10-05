package main

import (
	"compress/gzip"
	"os"
)

// 低レベルアクセスへの入り口（1）：io.Writer
// https://ascii.jp/elem/000/001/243/1243667/

func main() {
	file, err := os.Create("interface-io.writer-gzip.txt.gz")
	if err != nil {
		panic(err)
	}
	writer := gzip.NewWriter(file)
	writer.Write([]byte("gzip.Writer example\n"))
	writer.Close()
}
