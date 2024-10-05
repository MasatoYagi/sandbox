// https://go-tour-jp.appspot.com/concurrency/4

package main

import "fmt"

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, y+x
	}
	close(c) // これ以上の送信する値がないことを示す
}

func main() {
	c := make(chan int, 10)
	go fibonacci(cap(c), c)
	for i := range c { // チャネルが閉じられるまで、チャネルから値を繰り返し受信し続ける。chanのイテレーション時の返り値は1つだけ
		fmt.Println(i)
	}
}
