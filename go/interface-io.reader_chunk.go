package main

import (
	"fmt"
	"io"
	"os"
)

// 低レベルアクセスへの入り口（2）：io.Reader前編
// https://ascii.jp/elem/000/001/252/1252961/

func main() {
	for {
		buffer := make([]byte, 5)
		size, err := os.Stdin.Read(buffer)
		if err == io.EOF {
			fmt.Println("EOF")
			break
		}
		fmt.Printf("size=%d input='%s'\n", size, string(buffer))
	}
}

// $ go run stdin.go < stdin.go で5バイトずつ出力される
