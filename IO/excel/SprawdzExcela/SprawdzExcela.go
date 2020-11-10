package main

import (
	"errors"
	"log"
	"strconv"
	"strings"

	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/marek2222/goTest/IO/excel/SprawdzExcela/hlp"
)

// SprawdzExcela sprawdza pliki excel
func SprawdzExcela(rozszerz, katalog string) error {

	pliki, err := hlp.ListOfFilesSortByName(rozszerz, katalog)
	if err != nil {
		return err
	}

	for _, nazwa := range pliki {
		err := sprawdzDanyPlik(nazwa)
		if err != nil {
			return errors.New("błąd dla pliku '" + nazwa + "': " + err.Error())
		}
	}
	return nil
}

func sprawdzDanyPlik(nazwa string) error {

	p, err := excelize.OpenFile(nazwa)
	if err != nil {
		return errors.New("przy otwieraniu pliku: " + err.Error())
	}

	for index, sh := range p.GetSheetMap() {
		wiersze := p.GetRows(sh)
		log.Println("Arkusz: index:", index, ";  nazwa:", sh)
		log.Println("	Liczba kolumn i wierszy:", len(wiersze[0]), len(wiersze))
		log.Println("")

		sprawdzCzyPustyWiersz(p, sh, nazwa, wiersze)
		if err != nil {
			return errors.New("sprawdzCzyPustyWiersz(): " + err.Error())
		}

		sprawdzCzyZamowienieID(p, sh, nazwa, wiersze)
		if err != nil {
			return errors.New("sprawdzCzyZamowienieID(): " + err.Error())
		}
	}

	err = p.SaveAs(nazwa)
	if err != nil {
		return errors.New("xslx.SaveAs(): " + err.Error())
	}

	return nil
}

func sprawdzCzyZamowienieID(p *excelize.File, sh, nazwa string, wiersze [][]string) error {
	// B1 ZamowienieId
	// index = 2-1 = 1
	nrKol := 1
	slowo := strings.ToLower("ZamowienieId")
	czyZamID := false

	for nrWiersza, wiersz := range wiersze {

		zawieraZamID := strings.Contains(strings.ToLower(wiersz[nrKol]), slowo)
		if nrWiersza == 0 && zawieraZamID {
			czyZamID = true
		}

		if czyZamID {
			wartosc := wiersz[nrKol]
			zamID, err := strconv.Atoi(wartosc)
			if err != nil {
				return errors.New("strconv.Atoi(" + wartosc + "): " + err.Error())
			}

			if zamID > 0 {

			}
			//czyPustaKomorka()
		}

	}

	return nil
}

func sprawdzCzyPustyWiersz(p *excelize.File, sh, nazwa string, wiersze [][]string) error {

	liczbaKolumn := len(wiersze[0])

	for nrWiersza, wiersz := range wiersze {
		if czyPystyWiersz(wiersz, liczbaKolumn) {
			log.Println("Wiersz " + strconv.Itoa(nrWiersza+1) + " jest pusty i będzie usuwany!!!")
			p.RemoveRow(sh, nrWiersza)

			err := p.SaveAs(nazwa)
			if err != nil {
				return errors.New("xslx.SaveAs()2: " + err.Error())
			}

			err = sprawdzDanyPlik(nazwa)
			if err != nil {
				return errors.New("sprawdzDanyPlik()2: " + err.Error())
			}
		}
	}
	return nil
}

// sprawdza czy jest pusty wiersz
func czyPystyWiersz(wiersz []string, liczbaKolumn int) bool {
	czyPystyWiersz := true
	var lista []string
	kol := liczbaKolumn
	if kol > 10 {
		kol = 10
	}

	for _, komorka := range wiersz {
		lista = append(lista, komorka)
	}

	wartosc := strings.Join(lista, "")
	if len(wartosc) > 0 {
		czyPystyWiersz = false
	}

	return czyPystyWiersz
}
