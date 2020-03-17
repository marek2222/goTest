package APackages

import (
	"fmt"
	"math"
	"math/rand"
)

// Packages ...
func Packages() {
	var message string
	message = fmt.Sprintf("My favorite number is %d", rand.Intn(10))
	PrintMsg("Packages", message)

	message = fmt.Sprintf("Now you have %g problems.", math.Sqrt(7))
	PrintMsg("Import from math/rand", message)

	message = fmt.Sprintf("%f", math.Pi)
	PrintMsg("Exported names", message)
}

// PrintMsg show message with title
func PrintMsg(title string, message string) {
	fmt.Println(title, ":    ", message)
}
