package main

import "fmt"

func main() {
	// 型によって処理をスイッチ
	var i interface{}
	i = 100
	switch v := i.(type) {
	case int:
		fmt.Println(v * 2)
	case string:
		fmt.Println("文字列:" + v)
	default:
		fmt.Println("default")
	}
}
