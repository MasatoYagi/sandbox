package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

// https://pkg.go.dev/net/http#Handle
type countHandler struct {
	mu sync.Mutex
	n  int
}

// わざわざServeHTTPを定義する必要がある
// HandleFuncを使うと、型キャストにより定義の必要がなくなる
func (h *countHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.n++
	fmt.Fprintf(w, "count is %d\n", h.n)
}

func main() {
	// HTTPハンドラの作成
	http.Handle("/count", &countHandler{})
	// サーバーを起動
	// 第2引数でHTTPハンドラを指定
	// 基本的にGoのHTTPサーバーには1つのハンドラしか登録できない
	// (nilで省略した場合に、http.HandleFuncなどで登録したハンドラが利用される=>内部的にマルチプレクサ(http.ServerMux(デフォルトのhttp.DefaultServerMux))を利用して、パスによって利用するハンドラを切り替えている)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// # HTTPハンドラ
// 引数にレスポンスを書き込む先とリクエストを取る。
// 第1引数はレスポンスの書き込み先
// ・書き込みにはfmtパッケージの関数などが使える。
// 第2引数はクライアントからのリクエスト
// func handler(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprint(w, "Hello, HTTPサーバ")
// }

// HTTPハンドラはインターフェースとして定義されている
// ・ServeHTTPメソッドを実装していればハンドラとして扱われる
// type Handler interface {
// 	ServeHTTP(ResponseWriter, *Request)
// }

// 関数にインターフェースを実装できると、無名関数をキャストするだけでインターフェースを実装することができるのが有効な点
// net/httpパッケージで、関数にインターフェースを実装して活用している例
// https://cs.opensource.google/go/go/+/refs/tags/go1.23.1:src/net/http/server.go;l=2764-2773
// func (mux *ServeMux) HandleFunc(pattern string, handler func(ResponseWriter, *Request)) { // Handler型でない1️⃣
// 	if use121 {
// 		mux.mux121.handleFunc(pattern, handler)
// 	} else {
// 		mux.register(pattern, HandlerFunc(handler)) // 引数で受け取った関数をHandlerFunc型にキャストしている(ServeHTTPを実装する必要がない)2️⃣
// 	}
// }

// https://cs.opensource.google/go/go/+/refs/tags/go1.23.1:src/net/http/server.go;l=2795-2799
// ↑のHandleFunc()のラッパー
// func (mux *ServeMux) register(pattern string, handler Handler) { // ハンドラを実装するには、メソッド(ServeHTTP)を定義しておく必要がある3️⃣
// 	if err := mux.registerErr(pattern, handler); err != nil {
// 		panic(err)
// 	}
// }

// https://cs.opensource.google/go/go/+/refs/tags/go1.23.1:src/net/http/server.go;l=2216
// type HandlerFunc func(ResponseWriter, *Request) // 関数をベースにした型

// ServeHTTP calls f(w, r).
// func (f HandlerFunc) ServeHTTP(w ResponseWriter, r *Request) { // メソッドを持つ=>http.Handlerインターフェースを満たす4️⃣
// 	f(w, r)
// }

// https://cs.opensource.google/go/go/+/refs/tags/go1.23.1:src/net/http/server.go;l=88-90
// type Handler interface {
// 	ServeHTTP(ResponseWriter, *Request)
// }
