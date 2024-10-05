// https://go-tour-jp.appspot.com/concurrency/5
// selectステートメントの例としては分かりづらい
package main

import "fmt"

func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x: // ここが直感的じゃない気がする
			x, y = y, y+x
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func main() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c) // fibonacciを10回呼び出している？
		}
		quit <- 0 // fibonacciにquit経由で0を送信
	}()
	fibonacci(c, quit) // cに送信、quitから受信
}
