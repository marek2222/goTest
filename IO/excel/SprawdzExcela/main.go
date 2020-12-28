package main

import (
	"log"

	"github.com/marek2222/goTest/IO/excel/SprawdzExcela/excel"
)

func main() {
	katKonwersja := "./files"

	err := excel.SprawdzExcela(".xlsx", katKonwersja)
	if err != nil {
		log.Fatal("Błąd SprawdzExcela(): ", err)
	}

}
