package main

import (
	"errors"
	"log"
	"strconv"
	"strings"

	"github.com/360EntSecGroup-Skylar/excelize/v2"
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
		wiersze, err := p.GetRows(sh)
		if err != nil {
			return errors.New("p.GetRows(sh): " + err.Error())
		}
		log.Println("Arkusz: index:", index, ";  nazwa:", sh)
		log.Println("	Liczba kolumn i wierszy:", len(wiersze[0]), len(wiersze))
		log.Println("")

		err = usunPustyWiersz(p, sh, wiersze, nazwa)
		if err != nil {
			return errors.New("usunPustyWiersz(): " + err.Error())
		}

		err = sprawdzKolInt(p, wiersze, sh, nazwa, "ZamId", true)
		if err != nil {
			return errors.New("sprawdzKolInt(): " + err.Error())
		}
	}

	err = p.SaveAs(nazwa)
	if err != nil {
		return errors.New("xslx.SaveAs(): " + err.Error())
	}

	return nil
}

func sprawdzKolInt(p *excelize.File, wiersze [][]string, sh, nazwa, nazwaKol string, wymagana bool) error {
	idxKol := -1
	nazwaKol = strings.ToLower(nazwaKol)

	for nrWiersza, wiersz := range wiersze {

		// w pierwszym wierszu szukaj idxKol
		if idxKol == -1 {
			for idx := 0; idx < len(wiersz); idx++ {
				czyNazwaKol := strings.Contains(strings.ToLower(wiersz[idx]), nazwaKol)
				if nrWiersza == 0 && czyNazwaKol {
					idxKol = idx
					break
				}
			}
		} else {

			wartStr := wiersz[idxKol]
			if wymagana {
				wartosc, err := strconv.Atoi(wartStr)
				if err != nil {
					return errors.New("strconv.Atoi(" + wartStr + "): " + err.Error())
				}

				if wartosc < 1 {
					return errors.New("wartość jest wymagana i musi być większa od zera: " + strconv.Itoa(wartosc))
				}
			} else {

				_, err := strconv.Atoi(wartStr)
				if err != nil {
					if wartStr == "" {
						os, err := excelize.CoordinatesToCellName(idxKol+1, nrWiersza+1)
						if err != nil {
							return errors.New("hlp.CoordinatesToCellName(): " + err.Error())
						}
						p.SetCellInt(sh, os, 0)
					} else {
						return errors.New("strconv.Atoi(" + wartStr + "): " + err.Error())
					}
				}
			}
		}
	}
	return nil
}

// func ustawOs(x, y int) string {
// 	wynik := ""
// 	litery := map[int]string{
// 		0:  "A",
// 		1:  "B",
// 		2:  "C",
// 		3:  "D",
// 		4:  "E",
// 		5:  "F",
// 		6:  "G",
// 		7:  "H",
// 		8:  "I",
// 		9:  "J",
// 		10: "K",
// 		11: "L",
// 		12: "M",
// 		13: "N",
// 		14: "O",
// 		15: "P",
// 		16: "Q",
// 		17: "R",
// 		18: "S",
// 		19: "T",
// 		20: "U",
// 		21: "V",
// 		22: "W",
// 		23: "X",
// 		24: "Y",
// 		25: "Z",
// 	}
// 	wynik = litery[x]
// 	return wynik
// }

func usunPustyWiersz(p *excelize.File, sh string, wiersze [][]string, nazwa string) error {

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
