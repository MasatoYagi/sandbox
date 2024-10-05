package main

import "fmt"

// 二分木探索
// メソッド内でメソッドを再帰的に呼び出す
func main() {
	var it *IntTree
	it = it.Insert(5)
	it = it.Insert(3)
	it = it.Insert(10)
	it = it.Insert(2)

	fmt.Println(it.Contains(2))
	fmt.Println(it.Contains(12))
}

type IntTree struct {
	val         int
	left, right *IntTree
}

func (it *IntTree) Insert(val int) *IntTree { // 新しいデータを追加
	if it == nil {
		return &IntTree{val: val}
	}
	if val < it.val {
		it.left = it.left.Insert(val) // 再帰的に呼び出す
	} else if val > it.val {
		it.right = it.right.Insert(val) // 再帰的に呼び出す
	}
	return it
}

func (it *IntTree) Contains(val int) bool {
	switch {
	case it == nil:
		return false
	case val < it.val:
		return it.left.Contains(val) // 再帰的に呼び出す
	case val > it.val:
		return it.right.Contains(val) // 再帰的に呼び出す
	default:
		return true
	}
}
