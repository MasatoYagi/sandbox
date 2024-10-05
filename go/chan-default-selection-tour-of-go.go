// https://go-tour-jp.appspot.com/concurrency/6
package main

import (
	"fmt"
	"time"
)

func main() {
	tick := time.Tick(100 * time.Millisecond)
	boom := time.After(500 * time.Millisecond)
	for {
		select {
		case v1 := <-tick:
			fmt.Println(v1)
			fmt.Println("tick.")
		case v2 := <-boom:
			fmt.Println(v2)
			fmt.Println("BOOM!")
			return // 500ms後に終了
		default:
			fmt.Println("    .")
			time.Sleep(50 * time.Millisecond)
		}
	}
}
