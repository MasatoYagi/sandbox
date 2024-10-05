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

type S2 struct {
	N int
}

type S20 struct {
	N  int
	S2 // 名前のないフィールドになる
}

func (s *S20) Do() {
	fmt.Println(s.N)
}

func main() {
	s := S20{S2: S2{N: 2}, N: 20}
	s.Do() //=> 20(継承でないので2ではない！！！！！)
}
