package main

import "fmt"

type ServiceA struct {
	b ServiceBInterface
}

func (a ServiceA) ProcessingA() string {
	res := a.b.ProcessingB() // ServiceBに依存しているため、ServiceAの単体テストが実行できない
	if res == 0 {
		return "hello"
	} else {
		return "world"
	}
}

type ServiceB struct{}

type ServiceBInterface interface {
	ProcessingB() int
}

func (b ServiceB) ProcessingB() int {
	return 0
}

func main() {
	a := ServiceA{
		b: ServiceB{}, // 実行コード(not テスト)では、ServiceBのインスタンスを渡す
	}
	result := a.ProcessingA()
	fmt.Println(result)
}
