package main

import "fmt"

type Func1 func() string

func (f Func1) String() string { return f() }

func main() {
	// インターフェースを満たしているかチェックするイディオム
	var _ fmt.Stringer = Func1(nil) // Stringerインターフェースを満たしていなければ、コンパイルエラーが発生する
	// =>Func1 does not implement fmt.Stringer (missing method String)
}
