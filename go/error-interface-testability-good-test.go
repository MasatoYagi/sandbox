package main

import "fmt"

type serviceA struct {
	b ServiceBInterface
}

func (a serviceA) ProcessingA() string {
	res := a.b.ProcessingB() // ServiceBに依存しているため、ServiceAの単体テストが実行できない
	if res == 0 {
		return "hello"
	} else {
		return "world"
	}
}

type serviceB struct{}

type serviceBInterface interface {
	ProcessingB() int
}

func (b serviceB) ProcessingB() int {
	return 0
}

type mockServiceB struct {
	ProcessingBFunc func() int // 無名関数？
}

func (m *mockServiceB) ProcessingB() int {
	return m.ProcessingBFunc()
}

func main() {
	mockB := &mockServiceB{
		ProcessingBFunc: func() int { return 69 }, // 無名関数で渡せる（なんか定義しなくてよい）
	}

	a := ServiceA{b: mockB}

	result := a.ProcessingA()

	fmt.Println(result)
}
