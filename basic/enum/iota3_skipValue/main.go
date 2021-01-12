package main

import "fmt"

const (
	C1 = iota + 1
	_
	C3
	C4
)

func main() {
	fmt.Println(C1, C3, C4) // "1 3 4"
}
