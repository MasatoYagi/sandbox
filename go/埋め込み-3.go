package main

import "fmt"

// https://youtu.be/RAo0hEZd0fk?list=PL9MOSAifWs3xh--p7w8qPMXqTMuX0bBUn&t=46
// 埋め込みは継承でない！！！

// ■型リテラルでなければ埋め込められる
//  (型リテラル(type literal): []byte や *int など型の名前と記号を組み合わせて表現されている名前のない型)
//  - typeで定義したものや組み込み型
//  - インターフェースも埋められる
// ■インターフェースの実装
//  - 埋め込んだ値のメソッドもカウント

// Stringerを実装
type Hex1 int

func (h Hex1) String() string {
	return fmt.Sprintf("%x", int(h))
}

// Hex2もStringerを実装
type Hex10 struct{ Hex1 }

func main() {
	// Hex10(埋め込み先)はHex1(String()メソッドが定義され,Stringerインターフェースを満たす)が埋め込まれているので、
	// Stringerとして利用できる
	var s fmt.Stringer = Hex10{Hex1: 1}
	fmt.Println(s)
}
