package main

import (
	"fmt"
	"sync"
)

func main() {
	var count int
	var lock sync.Mutex

	increment := func() {
		lock.Lock()
		defer lock.Unlock()
		count++
		fmt.Printf("incrementing: %d\n", count)
	}

	decrement := func() {
		lock.Lock()
		defer lock.Unlock()
		count--
		fmt.Printf("decrementing: %d\n", count)
	}

	var artimentic sync.WaitGroup
	for i := 0; i < 5; i++ {
		artimentic.Add(1)
		go func() {
			defer artimentic.Done()
			increment()
		}()
	}
	for i := 0; i < 5; i++ {
		artimentic.Add(1)
		go func() {
			defer artimentic.Done()
			decrement()
		}()
	}
	artimentic.Wait()
}
