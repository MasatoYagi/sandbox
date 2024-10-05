package main

import (
	"fmt"
	"strings"
)

// https://ascii.jp/elem/000/001/260/1260449/2/

// データ型を指定して解析
// io.Reader から読み込んだデータは、今のところ単なるバイト列か文字列としてしか扱っていませんでした。
// io.Reader のデータを整数や浮動少数点数に変換するには、 fmt.Fscan を使います。
// 1つめの引数に io.Reader を渡し、それ以降に変数のポインタを渡すと、その変数にデータが書き込まれます。

// fmt.Fscan はデータがスペース区切りであることを前提としています。 fmt.Fscanln は改行区切り時に使います。
var source2 = "123 1.234 1.0e4 test"

func main() {
	reader := strings.NewReader(source2)
	var i int
	var f, g float64
	var s string
	// fmt.Fscanf を使うと任意のデータ区切りをフォーマット文字列として指定できます。
	fmt.Fscan(reader, &i, &f, &g, &s)
	fmt.Printf("i=%#v f=%#v g=%#v s=%#v\n", i, f, g)
	// 似た名前のC言語の関数をご存知の方もいると思いますが、Go言語は型情報をデータが持っているため、
	// すべて「%v」と書いておけば変数の型を読み取って変換してくれます。
}
