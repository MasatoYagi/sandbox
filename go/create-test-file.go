package main

import (
	"archive/zip"
	"crypto/rand"
	"io"
	"os"
	"strings"
)

// 低レベルアクセスへの入り口（2）：io.Reader前編
// https://ascii.jp/elem/000/001/252/1252961/

func main() {
	// Q2. テスト用の適当なサイズのファイルを作成
	bf := make([]byte, 1024)
	_, _ = rand.Read(bf)
	file2, _ := os.OpenFile("create-test-file.txt", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666) // os.Create()と同じ動作
	defer file2.Close()
	file2.Write(bf)
	// 模範解答
	file2_1, _ := os.OpenFile("create-test-file.txt", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666) // os.Create()と同じ動作
	defer file2_1.Close()
	io.CopyN(file2_1, rand.Reader, 1024)

	// Q3. zipファイルの書き込み
	// 単なるファイルのzip圧縮
	file3, _ := os.Create("create-test-file.zip")
	defer file3.Close()
	zipWriter := zip.NewWriter(file3) // zipファイルの書き込み用の構造体(io.Writerは満たさない)
	defer zipWriter.Close()
	_, _ = zipWriter.Create("create-test-file.txt") // Createでio.Writerが返される&file3(zipファイル)にtxtファイルを圧縮したデータが格納される
	// A3. 模範解答：Create()を使ってio.Writerを作る
	// zipの内容を書き込むファイル
	file3_1, _ := os.Create("create-test-file3-1.zip")
	defer file3_1.Close()
	// zipファイル
	zipWriter3_1 := zip.NewWriter(file3_1)
	defer zipWriter3_1.Close()
	// ファイルの数だけ書き込み
	a, _ := zipWriter3_1.Create("a.txt") // *zip.Writerはio.Writerではないが、Create()メソッドを呼ぶと個別のファイルを書き込むためのio.Writerが返ってくる
	io.Copy(a, strings.NewReader("1つめのファイルのテキストです"))
	b, _ := zipWriter3_1.Create("b.txt") // *zip.Writerはio.Writerではないが、Create()メソッドを呼ぶと個別のファイルを書き込むためのio.Writerが返ってくる
	io.Copy(b, strings.NewReader("2つめのファイルのテキストです"))
}

// $ go run stdin.go < stdin.go で5バイトずつ出力される
