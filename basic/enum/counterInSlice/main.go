package main

import (
	"fmt"

	entities "github.com/marek2222/goTest/basic/enum/counterInSlice/entities"
)

func main() {
	var products = []entities.Product{
		{
			ID:       "p01",
			Name:     "tivi 1",
			Price:    5,
			Quantity: 9,
			Status:   false,
		},
		{
			ID:       "p02",
			Name:     "tivi 2",
			Price:    2,
			Quantity: 8,
			Status:   true,
		},
		{
			ID:       "p03",
			Name:     "laptop 3",
			Price:    11,
			Quantity: 7,
			Status:   false,
		},
	}

	var min float64 = 5
	var max float64 = 20
	result := Count(min, max, products)
	fmt.Println("Result: ", result)
}

// Count ...
func Count(min, max float64, products []entities.Product) (counter int) {
	counter = 0
	for _, product := range products {
		if product.Price >= min && product.Price <= max {
			counter++
		}
	}
	return counter
}
