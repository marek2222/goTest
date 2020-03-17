package APackages

import "fmt"

// add two int value
func add(x int, y int) int {
	return (x + y)
}

// FunctionAdd ...
func FunctionAdd() {
	fmt.Println("Function add 42 and 13: ", add(42, 13))
}
