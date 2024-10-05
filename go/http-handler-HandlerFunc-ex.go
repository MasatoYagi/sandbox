package main

import (
	"fmt"
	"io"
	"net/http"
)

// https://youtu.be/ExHAphUgHgY?list=PL9MOSAifWs3xh--p7w8qPMXqTMuX0bBUn&t=2077
func main() {
	// HTTPハンドラの作成
	// 引数の無名関数は、HandlerFunc内でキャスト(↓のような感じでServeHTTPメソッドとして関数が実行されている;関数にインターフェースを実装するメリット)
	// => func ServeHTTP(w http.ResponseWriter, _ *http.Request) { return f() }
	h1 := http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) { // HandlerFunc型にキャストすることで、HandlerFunc型のServeHTTPメソッドを付与できる!!!!!
		io.WriteString(w, "Hello from a HandleFunc #1!\n")
	})
	fmt.Printf("%T\n", h1) // http.HandlerFunc
	http.HandleFunc("/", h1)
	http.ListenAndServe(":8080", h1)
}
