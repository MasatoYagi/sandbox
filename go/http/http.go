package main

import (
	"log"
	"net/http"
)

func main() {
	http.Handle("/", &templateHandler{filename: "chat.html"})
	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
