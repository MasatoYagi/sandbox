package main

import "fmt"

func main() {
	ch1 := make(chan int)
	ch2 := make(chan string)
	go func() { ch1 <- 100 }()
	go func() { ch2 <- "hi" }()

	select { // 先にchannelに入った方のデータが標準出力される
	case v1 := <-ch1:
		fmt.Println(v1)
	case v2 := <-ch2:
		fmt.Println(v2)
	}
}
