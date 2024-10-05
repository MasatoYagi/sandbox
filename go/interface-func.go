// 関数にインターフェースを実装させる
// https://youtu.be/ExHAphUgHgY?list=PL9MOSAifWs3xh--p7w8qPMXqTMuX0bBUn&t=1075
package main

import "fmt"

// Func型のデータ型を作成
type Func func() string

// FuncにString()メソッドを追加
func (f Func) String() string { return f() } // 無名関数をキャストできるようになる

func main() {
	// Func側へのキャスト
	var s fmt.Stringer = Func(func() string { // 無名関数 Goでは関数も第一級オブジェクト（intやstringと同じ）なので、基底型にして新しく型を作ることが可能
		return "hi"
	})
	fmt.Println(s)
}

// net/http
// https://cs.opensource.google/go/go/+/refs/tags/go1.23.1:src/net/http/server.go;l=2764-2773
