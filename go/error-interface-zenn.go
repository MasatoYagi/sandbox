// https://zenn.dev/kasa/articles/golang-interface#%E3%82%BD%E3%83%BC%E3%82%B9%E3%82%B3%E3%83%BC%E3%83%89%E3%81%AE%E6%8B%A1%E5%BC%B5%E6%80%A7%E5%90%91%E4%B8%8A
package main

import "fmt"

type MyError struct {
	Code    int
	Message string
}

func (e MyError) Error() string {
	return fmt.Sprintf("error: code=%d, message=%s", e.Code, e.Message)
}

func (e MyError) IsCritical() bool {
	return e.Code >= 100
}

func causeError() error {
	return MyError{
		Message: "An error occurred",
		Code:    1,
	}
}

func main() {
	err := causeError()
	if err != nil {
		fmt.Println(err)

		if myErr, ok := err.(MyError); ok {
			if myErr.IsCritical() {
				fmt.Println("Critical error encountered!")
			} else {
				fmt.Println("Non-critical error encounterd.")
			}
		}
	}
}
