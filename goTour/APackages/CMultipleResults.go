package APackages

import (
	"fmt"
)

func swap(x, y string) (string, string) {
	return y, x
}

// CMultipleResults ...
func CMultipleResults() {
	fmt.Println()
	fmt.Println("MultipleResults:")

	a, b := swap("hello", "world")
	fmt.Println("  ", a, b)

}
