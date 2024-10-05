package main

import "fmt"

func main() {
	byteArray := []byte("ASCII")
	fmt.Println(byteArray)                              // byteArrayは[]byte{0x41, 0x53, 0x43, 0x49, 0x49}
	str := string([]byte{0x41, 0x53, 0x43, 0x49, 0x49}) // string型(UTF-8形式)
	fmt.Println(str)                                    // strは"ASCII"
}
