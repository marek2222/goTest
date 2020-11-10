package main

import (
	"log"
)

func main() {
	katKonwersja := "./files"

	err := SprawdzExcela(".xlsx", katKonwersja)
	if err != nil {
		log.Fatal("Błąd SprawdzExcela(): ", err)
	}

}
