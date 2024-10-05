package main

import (
	"fmt"
	"io"
	"os"
)

// 低レベルアクセスへの入り口（2）：io.Reader前編
// https://ascii.jp/elem/000/001/252/1252961/

func main() {
	// 標準入力で受け取り、標準出力するだけの処理
	fmt.Println("標準入力")
	buffer := make([]byte, 5)       // 読み取り時はバッファを確保する必要がある
	_, err := os.Stdin.Read(buffer) /// 読み込みの際はバッファを利用する
	if err == io.EOF {
		fmt.Println("オーバー")
		return
	}
	fmt.Println("標準出力")
	os.Stdout.Write(buffer) // バッファの値を書き込み

	// 全て読み込む
	file0, _ := os.Open("interface-io.reader-0.txt")
	buffer, _ = io.ReadAll(file0) // メモリに収まるサイズであれば、これを利用することが多いハズ

	// ファイルを読み取り、再度ファイルに書き込むだけの処理
	// `os.Create()` 関数は、指定したファイルが存在しない場合は新しいファイルを作成し、存在する場合はファイルを空にして再作成する
	file1, _ := os.Open("interface-io.reader-1.txt")
	defer file1.Close()
	file2, _ := os.Create("interface-io.reader-2.txt")
	defer file2.Close()
	_, _ = io.Copy(file2, file1) // 確保するバッファは32KB

	// 読み取り時に確保するバッファサイズを指定
	file3, _ := os.Open("interface-io.reader-3.txt")
	defer file3.Close()
	file4, _ := os.Create("interface-io.reader-4.txt")
	defer file4.Close()
	_, _ = io.CopyN(file4, file3, 3) // 読み取り時のバッファサイズが3byte分だけなので、最初の1文字しかコピーされない
}
