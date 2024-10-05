package main

import (
	"fmt"
	"net/http"
)

type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error) // http.Clientに定義されているメソッド
}

func main() {
	var cli *http.Client = http.DefaultClient
	var cli2 HTTPClient = cli // Goでは後からinterfaceを実装できる(Javaだとできない)
	fmt.Println(cli2)
}
