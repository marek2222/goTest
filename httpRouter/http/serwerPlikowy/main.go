package main

import (
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("c:/_Projects/go/src/github.com/marek2222/goTest/")))
	http.ListenAndServe(":8081", nil)
}
