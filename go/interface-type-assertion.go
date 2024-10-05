package main

import "fmt"

func main() {
	// 型アサーション
	// インターフェース.(型)
	// - インターフェース型の値を任意の型にキャストする
	// - 第2戻り値にキャストできるかどうかが返る
	var v interface{}
	v = 100
	n, ok := v.(int)
	fmt.Println(n, ok)

	s, ok := v.(string)
	fmt.Println(s, ok)
}
