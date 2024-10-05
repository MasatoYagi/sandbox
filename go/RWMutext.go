package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var m sync.RWMutex
	m.RLock()
	go func() {
		time.Sleep(3 * time.Second)
		m.RUnlock()
		fmt.Println("unlock 1")
	}()
	m.RLock()
	m.RUnlock()
	fmt.Println("unlock 2")
}
