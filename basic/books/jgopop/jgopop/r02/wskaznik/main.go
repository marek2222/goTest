package main

import "fmt"

func main() {
	// aktualizacja pośrednia
	//     przez wskaźnik p wskazujący na x
	aktualizacjaPosrednia()

	// porównanie wskaźników
	porownanieWskazników()

	// zwracanie adresu zmiennej lokalnej jest bezpieczne dla funkcji
	var p = f()
	fmt.Println(p)
	fmt.Println(f() == f())
}

func aktualizacjaPosrednia() {
	x := 1
	p := &x // wskażnik do zmiennej, p wskazuje na x lub p zawiera adres x
	fmt.Println(*p)

	*p = 2 // to przypisanie aktualizuje zmienna przez wskaźnik p wskazujący na x
	fmt.Println(*p)
}

func porownanieWskazników() {
	var x, y int
	fmt.Println(&x == &x, &x == &y, &x == nil) // powinno być "true false false"
}

// zwracanie adresu zmiennej lokalnej
func f() *int {
	v := 1
	return &v
}
