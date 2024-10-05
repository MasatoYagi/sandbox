// https://youtu.be/ExHAphUgHgY?list=PL9MOSAifWs3xh--p7w8qPMXqTMuX0bBUn&t=695
package main

import "fmt"

func main() {
	var s Stringer = Hex(100)
	fmt.Println(s)
	fmt.Println(s.String()) // Hexを直接利用すると、Hexに変更が加わった時に修正(?)が必要になる
}

type Hex int // int型のHex型を作成

func (h Hex) String() string { // Hex型にStringメソッドを定義(String型を満たす)
	return fmt.Sprintf("%x", int(h))
}

type Stringer interface { // Stringerインターフェース(Stringメソッドを持つデータ)を定義
	String() string
}
