// https://www.youtube.com/watch?v=q0vk1nrT96I&list=PL9MOSAifWs3whvWOsObk3uCBXtVAhD2A7&index=28
package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int) // 容量0(make(chan int 0))

	go func() {
		ch <- 100 // 送信
	}()

	go func() {
		v := <-ch // 受信
		fmt.Println(v)
	}()
	time.Sleep(1 * time.Second)
}
