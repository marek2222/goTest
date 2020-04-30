package main

import (
	"fmt"

	"github.com/goTest/goTour/APackages"
)

// Wycieczka ...
type Wycieczka struct {
	HotelID     int
	WycieczkaID int
}

func main() {
	// Packages
	APackages.APackages()
	APackages.BFunctionAdd()
	APackages.CMultipleResults()


	
	// wycieczki := []Wycieczka{}
	// wycieczki = append(wycieczki, Wycieczka{1, 5})
	// wycieczki = append(wycieczki, Wycieczka{3, 6})
	// wycieczki = append(wycieczki, Wycieczka{9, 5})
	// wycieczki = append(wycieczki, Wycieczka{4, 2})

	// fmt.Println(wycieczki)
	// unikalneWcID := unikalneWcID(wycieczki)
	// fmt.Println(unikalneWcID)
}

func unikalneWcID(wycieczki []Wycieczka) (wyniki []Wycieczka) {
	klucze := make(map[int][]Wycieczka)
	for i, encja := range wycieczki {
		if _, wartosc := klucze[encja.WycieczkaID]; !wartosc {
			klucze[wycieczki[i].WycieczkaID] = append(klucze[wycieczki[i].WycieczkaID], wycieczki[i])
			wyniki = append(wyniki, wycieczki[i])
		}
	}
	return wyniki
}
