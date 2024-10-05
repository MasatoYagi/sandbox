package main

import (
	"io"
	"os"
	"strings"
)

func main() {
	file, _ := os.Open("バイナリ解析.txt")
	lReader := io.LimitReader(file, 16)
	io.Copy(os.Stdout, lReader) // たくさんデータがあっても先頭の16バイトしか読み込めないようにする。

	reader := strings.NewReader("Example of io.SectionReader\n")
	sectionReader := io.NewSectionReader(reader, 14, 7) // =>SectionReader
	io.Copy(os.Stdout, sectionReader)
}
