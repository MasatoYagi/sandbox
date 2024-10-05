package main

import (
	"bytes"
	"fmt"
)

func main() {
	buf1 := new(bytes.Buffer)
	fmt.Printf("%T\n", buf1)

	buf2 := new(bytes.Buffer)
	fmt.Printf("%T\n", buf2)
}
