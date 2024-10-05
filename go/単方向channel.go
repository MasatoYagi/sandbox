// https://www.youtube.com/watch?v=q0vk1nrT96I&list=PL9MOSAifWs3whvWOsObk3uCBXtVAhD2A7&index=28
package main

import "fmt"

func main() {
	ch := makeCh()
	go func(ch chan<- int) { // 送信専用チャネル,
		ch <- 100
	}(ch) // 双方向チャネル→単方向チャネルに明示的にキャストする必要はない
	fmt.Println(recvCh(ch))
}

func makeCh() chan int {
	return make(chan int)
}

func recvCh(recv <-chan int) int { // 受信専用チャネル
	return <-recv // チャネルから値を取得
}
