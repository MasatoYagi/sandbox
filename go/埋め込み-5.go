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

// # 既存のインターフェースの振る舞いを変える
//
// 特定のメソッドだけをデコレーションしたい場合に使える
// https://tenntenn.dev/ja/posts/qiita-eac962a49c56b2b15ee8/
type S4 struct {
	Name string
}

func (s S4) M1() {
	fmt.Println(s.Name)
}

type S4Interface interface {
	M1()
}

type S40 struct{ S4Interface } // インターフェースを埋め込む

func (s S40) M1() { // 元のメソッドを呼ぶ(M1の振る舞いを変えて、一部の実装だけを変えている)
	fmt.Print("元のメソッドから変更: ") // 追加された振る舞い
	s.S4Interface.M1()        // 元のメソッド
}

func ToS40(s S4Interface) S4Interface {
	return S40{s}
}

func main() {
	s4 := S4{Name: "S4"}
	s4.M1() //=>S4
	s40 := ToS40(S4{Name: "ToS40"})
	s40.M1() //=>元のメソッドから変更: ToS40
}
