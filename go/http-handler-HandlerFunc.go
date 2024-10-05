package main

import (
	"io"
	"log"
	"net/http"
)

// !!!HTTPハンドラ=http.Handlerインターフェース!!!

func main() {
	// HTTPハンドラの作成
	h1 := func(w http.ResponseWriter, _ *http.Request) {
		io.WriteString(w, "Hello from a HandleFunc #1!\n")
	}
	h2 := func(w http.ResponseWriter, _ *http.Request) {
		io.WriteString(w, "Hello from a HandleFunc #2!\n")
	}
	// HTTPハンドラとURLパスを紐づけ(DefaultServeMux.HandleFunc()経由で、DefaultServeMuxのにハンドラを登録)
	http.HandleFunc("/", h1)         // h1, h2はHandlerFunc型にキャスト(+ServeHTTPメソッドが付与され,Handlerインターフェースを満たす)され,HTTPハンドラになる
	http.HandleFunc("/endpoint", h2) // ↑により、HTTPハンドラを作る時の手間(インターフェースを満たすためのメソッド定義)が不要になる
	// サーバーを起動
	// 第2引数でHTTPハンドラを指定
	// 基本的にGoのHTTPサーバーには1つのハンドラしか登録できない
	// nilで省略した場合に、http.HandleFuncなどで登録したハンドラが利用される=>内部的にマルチプレクサ(http.ServerMux(デフォルトのhttp.DefaultServeMux))を利用して、パスによって利用するハンドラを切り替えている
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

// DefaultServeMuxはServeMux型のハンドラで，最初からあるデフォルトのハンドラ(http.ListenAndServe()の第2引数がnilの時に用いられる)

// # Handler関連のメソッド・クラス

// ## 1. responseにメソッドを使う
// - HandleFunc(登録する関数(ハンドラ))
// - HandlerFunc(関数(ハンドラ)を登録) https://pkg.go.dev/net/http#HandlerFunc
// エンドポイントと関数を一対一で紐付け，DefaultServeMuxに関数として設定する
// h1 := func(w http.ResponseWriter, _ *http.Request) { // Goではhttpハンドラはfunc(w http.ResponseWriter, _ *http.Request)のシグネチャで定義する必要がある
// 	io.WriteString(w, "Hello from a HandleFunc #1!\n")
// }
// http.HandleFunc("/", h1) // 関数(h1)を利用

// 参考: https://qiita.com/immrshc/items/1d1c64d05f7e72e31a98#%E3%82%88%E3%81%8F%E5%87%BA%E3%81%A6%E3%81%8F%E3%82%8B%E5%9E%8B%E5%89%8D%E6%8F%90

// ## 2. responseに構造体(struct)を使う
// - Handle()
// - Handler()
// Handlerを使う場合，ServeHttpを構造体内部のメソッドに設定し，構造体を返してresponse作成を行う
// type countHandler struct {
// 	mu sync.Mutex
// 	n  int
// }
// func (h *countHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	h.mu.Lock()
// 	defer h.mu.Unlock()
// 	h.n++
// 	fmt.Fprintf(w, "count is %d\n", h.n)
// }
// func main() {
// 	http.Handle("/count", &countHandler{}) // 構造体(HTTPハンドラのインターフェースを満たす)
// 	log.Fatal(http.ListenAndServe(":8080", nil))
// }

// https://qiita.com/immrshc/items/1d1c64d05f7e72e31a98#%E3%82%88%E3%81%8F%E5%87%BA%E3%81%A6%E3%81%8F%E3%82%8B%E5%9E%8B%E5%89%8D%E6%8F%90

// http.Handlerやhttp.HandlerFuncはDefaultServeMux.Handleを内部的に呼び出す糖衣構文
// https://golang.org/src/net/http/server.go?s=73173:73217#L2391
// func Handle(pattern string, handler Handler) { DefaultServeMux.Handle(pattern, handler) }
