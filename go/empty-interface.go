package main

import "fmt"

func main() {
	var v interface{}
	fmt.Println(v)

	v = 100
	fmt.Println(v)

	v = "hoge"
	fmt.Println(v)
}
