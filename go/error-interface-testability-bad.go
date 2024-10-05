package main

import "fmt"

type ServiceA struct{}

func (a ServiceA) ProcessingA() string {
	b := ServiceB{}
	res := b.ProcessingB() // ServiceBに依存しているため、ServiceAの単体テストが実行できない
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
	a := ServiceA{}
	result := a.ProcessingA()
	fmt.Println(result)
}
