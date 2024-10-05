package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

// func StringToUpper(s *strings.Reader) {
// 	data := make([]byte, 300)
// 	len, _ := s.Read(data)
// 	str := string(data[:len])

// 	result := strings.ToUpper(str)
// 	fmt.Println(result)
// }

// func FileToUpper(f *os.File) {
// 	data := make([]byte, 300)
// 	len, _ := f.Read(data)
// 	str := string(data[:len])

// 	result := strings.ToUpper(str)
// 	fmt.Println(result)
// }

func ToUpper(r io.Reader) { // 引数をio.Readerにすることで、strings.Readerとos.Fileを受け取れるようになる
	data := make([]byte, 300)
	len, _ := r.Read(data)
	str := string(data[:len])

	result := strings.ToUpper(str)
	fmt.Println(result)
}

func main() {
	// 文字列リーダーからの読み取り
	strReader := strings.NewReader("This is a sample string.")
	ToUpper(strReader)

	// ファイルからの読み取り
	file, err := os.Open("example.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
	}
	defer file.Close()

	ToUpper(file)
}
