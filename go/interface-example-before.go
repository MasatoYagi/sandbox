package main

import (
	"fmt"
	"os"
	"strings"
)

func StringToUpper(s *strings.Reader) {
	data := make([]byte, 300)
	len, _ := s.Read(data)
	str := string(data[:len])

	result := strings.ToUpper(str)
	fmt.Println(result)
}

func FileToUpper(f *os.File) {
	data := make([]byte, 300)
	len, _ := f.Read(data)
	str := string(data[:len])

	result := strings.ToUpper(str)
	fmt.Println(result)
}

func main() {
	// 文字列リーダーからの読み取り
	strReader := strings.NewReader("This is a sample string.")
	StringToUpper(strReader)

	// ファイルからの読み取り
	file, err := os.Open("example.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
	}
	defer file.Close()

	FileToUpper(file)
}
