// https://gist.githubusercontent.com/kaipakartik/8120855/raw/16a2eaf5793b82d9b1f1b89b5f48b4a236bede10/tree.go
// 感想：むずい
package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	WalkRecursive(t, ch) // 2つのコードに分けることでチャンネルを閉じることができるので、任意の大きさの二分木に対して探索を行える
	close(ch)
}

func WalkRecursive(t *tree.Tree, ch chan int) {
	if t != nil {
		WalkRecursive(t.Left, ch)
		ch <- t.Value
		WalkRecursive(t.Right, ch)
	}
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1, ch2 := make(chan int), make(chan int)
	go Walk(t1, ch1)
	go Walk(t2, ch2)
	for {
		n1, ok1 := <-ch1
		n2, ok2 := <-ch2
		if ok1 != ok2 || n1 != n2 {
			return false
		}
		if !ok1 {
			break
		}
	}
	return true
}

func main() {
	ch := make(chan int)
	go Walk(tree.New(1), ch)
	fmt.Println(Same(tree.New(1), tree.New(2)))
	fmt.Println(Same(tree.New(1), tree.New(1)))
	fmt.Println(Same(tree.New(2), tree.New(1)))
}

// // Walk walks the tree t sending all values
// // from the tree to the channel ch.
// func Walk(t *tree.Tree, ch chan int) {
// 	if t == nil {
// 		return
// 	}
// 	if t.Left != nil {
// 		Walk(t.Left, ch) // 再帰的に呼び出して、左側からチャネルに送信
// 	}
// 	ch <- t.Value // 左側の値から順にチャネルに送信
// 	fmt.Println(t.Value)
// 	if t.Right != nil {
// 		Walk(t.Right, ch)
// 	}
// }

// // Same determines whether the trees
// // t1 and t2 contain the same values.
// func Same(t1, t2 *tree.Tree) bool {
// 	ch1, ch2 := make(chan int), make(chan int)
// 	go Walk(t1, ch1)
// 	go Walk(t2, ch2)
// 	b := true
// 	for i := 0; i < 10; i++ {
// 		v1 := <-ch1
// 		v2 := <-ch2
// 		if v1 != v2 {
// 			b = false
// 		}
// 	}
// 	return b
// }

// func main() {
// 	b := Same(tree.New(1), tree.New(1))
// 	if b {
// 		fmt.Println("same")
// 	} else {
// 		fmt.Println("not same")
// 	}
// }
