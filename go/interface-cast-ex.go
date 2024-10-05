package main

import "fmt"

type Hoge1 struct {
	field1 string
}

func (h Hoge1) なにかのメソッド() string {
	return h.field1
}

type Hoge1Interface interface {
	なにかのメソッド() string
}

type Hoge2 struct {
	field1 string
}

func main() {
	a2 := Hoge2{field1: "Hoge2ですよ(キャスト前)"}
	var a2Neo Hoge1Interface = Hoge1(a2) // Hoge1にキャストすることで、Hoge2がインターフェースを満たすためのメソッド定義が不要になる！！！！
	fmt.Println(a2Neo.なにかのメソッド())
}
