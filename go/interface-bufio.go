package main

import (
	"bufio"
	"os"
)

func main() {
	// # バッファ付き出力
	buffer := bufio.NewWriter(os.Stdout)
	buffer.WriteString("bufio.Writer ")
	buffer.Flush()
	buffer.WriteString("example\n")
	buffer.Flush()
}
