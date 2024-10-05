package main

import (
	"fmt"
	"math"
)

func main() {
	var s Shape = Circle{Radius: 5.5}
	fmt.Println(s.Area())
	s = Rectangle{Width: 5.5, Height: 5.5}
	fmt.Println(s.Area())
}

type Shape interface {
	Area() float64
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

type Rectangle struct {
	Width, Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// 「バイト列 b を書き込み、書き込んだバイト数 n と、エラーが起きた場合はそのエラー error を返す」という振る舞いは、
// 通常のファイルに限らず、さまざまなものに適用できそうですよね。
// そこで、同じ仕様のメソッドを持つ型を統一的に扱えると便利そうです。
// Go言語では、その場合に使える仕組みとして インタフェース という型が用意されています。
// https://ascii.jp/elem/000/001/243/1243667/
