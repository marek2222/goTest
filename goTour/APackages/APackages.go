package APackages

import (
	"fmt"
	"math"
	"math/rand"

	"github.com/goTest/goTour/PrintMsg"
)

// APackages ...
func APackages() {
	fmt.Println()
	fmt.Println("Packages:")

	var message string
	message = fmt.Sprintf("My favorite number is %d", rand.Intn(10))
	PrintMsg.Print("Packages", message)

	message = fmt.Sprintf("Now you have %g problems.", math.Sqrt(7))
	PrintMsg.Print("Import from math/rand", message)

	message = fmt.Sprintf("%f", math.Pi)
	PrintMsg.Print("Exported names", message)
}
