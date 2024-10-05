package main

import "fmt"

// https://youtu.be/RAo0hEZd0fk?list=PL9MOSAifWs3xh--p7w8qPMXqTMuX0bBUn&t=46
// 埋め込みは継承でない！！！

type S1 struct {
	N int
}

type S10 struct {
	S1 // 名前のないフィールドになる
}

func main() {
	s := S10{S1: S1{N: 100}}
	// 型名を指定してアクセスできる
	fmt.Println(s.S1.N) // 100
	// S1のフィールドにアクセスできる
	fmt.Println(s.N) // 100
}
