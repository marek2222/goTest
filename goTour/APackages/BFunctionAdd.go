package APackages

import (
	"fmt"

	"github.com/goTest/goTour/PrintMsg"
)

// add two int value
func add(x int, y int) int {
	return (x + y)
}

// add two int value - short
func add2(x, y int) int {
	return (x + y)
}

// FunctionAdd ...
func BFunctionAdd() {
	fmt.Println()
	fmt.Println("Functions:")

	message := fmt.Sprintf("Add 42 and 13: %d", add(42, 13))
	PrintMsg.Print("Function", message)

	message = fmt.Sprintf("Add 42 and 13: %d", add2(42, 13))
	PrintMsg.Print("Short function", message)
}
