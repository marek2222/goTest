// https://medium.com/better-programming/easy-guide-to-unit-testing-in-golang-4fc1e9d96679
// Easy Guide to Unit Testing in Golang

package main

import (
	"fmt"
)

func main() {

	op := GetKeyOperator()

	key := op.Generate(2, 3)

	a, b, _ := op.Degenerate(key)

	fmt.Printf("key=%v, a=%v, b=%v", key, a, b)

}
