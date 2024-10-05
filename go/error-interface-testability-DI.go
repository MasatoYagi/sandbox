package main

import "fmt"

type ServiceA struct {
	b ServiceB
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

func (b ServiceB) ProcessingB() int {
	return 0
}

func main() {
	a := ServiceA{
		b: ServiceB{}, // ServiceB型のインスタンスをServiceAの外部から渡せるように依存性の注入（Dependency Injection, DI）を行った。
	}
	result := a.ProcessingA()
	fmt.Println(result)
}
