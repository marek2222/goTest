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
		log.Println("Plik:", nazwa, " Arkusz:", index, " nazwa:", sh, " kolumn:", len(wiersze[0]), " wierszy:", len(wiersze))
		log.Println("")

		err = usunPustyWiersz(p, sh, wiersze, nazwa)
		if err != nil {
			return errors.New("usunPustyWiersz(): " + err.Error())
		}

		err = musiKolInt(p, wiersze, sh, nazwa, "ZamId")
		if err != nil {
			return errors.New("musiKolInt(): " + err.Error())
		}

		err = ustawKolInt(p, wiersze, sh, nazwa, "ZamId")
		if err != nil {
			return errors.New("musiKolInt(): " + err.Error())
		}
	}

	err = zapiszExcela(p, nazwa)
	if err != nil {
		return err
	}

	return nil
}

func musiKolInt(p *excelize.File, wiersze [][]string, sh, nazwa, nazwaKol string) error {
	var err error
	idxKol := -1
	nazwaKol = strings.ToLower(nazwaKol)

	for nrWiersza, wiersz := range wiersze {

		// w pierwszym wierszu szukaj idxKol, w kolejnych sprawdz wartości
		if idxKol == -1 {
			idxKol, err = indexKolPoNazwie(wiersz, nrWiersza, nazwaKol)
			if err != nil {
				return err
			}
		} else {
			wartStr := wiersz[idxKol]
			err := czyIntWiekszyOdZera(wartStr)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func ustawKolInt(p *excelize.File, wiersze [][]string, sh, nazwa, nazwaKol string) error {
	var err error
	idxKol := -1
	nazwaKol = strings.ToLower(nazwaKol)

	for nrWiersza, wiersz := range wiersze {

		// w pierwszym wierszu szukaj idxKol, w kolejnych sprawdz wartości
		if idxKol == -1 {
			idxKol, err = indexKolPoNazwie(wiersz, nrWiersza, nazwaKol)
			if err != nil {
				return err
			}
		} else {
			wartStr := wiersz[idxKol]
			nrKol := idxKol + 1
			nrWie := nrWiersza + 1
			err := ustawKomorkeJakoIntZero(p, wartStr, sh, nrKol, nrWie)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func ustawKomorkeJakoIntZero(p *excelize.File, wartStr, sh string, nrKol, nrWie int) error {
	_, err := strconv.Atoi(wartStr)
	if err != nil {
		return errors.New("strconv.Atoi(" + wartStr + "): " + err.Error())
	}

	if wartStr == "" {
		os, err := excelize.CoordinatesToCellName(nrKol, nrWie)
		if err != nil {
			return errors.New("hlp.CoordinatesToCellName(): " + err.Error())
		}

		p.SetCellInt(sh, os, 0)
	}
	return nil
}

func czyIntWiekszyOdZera(wartStr string) error {
	wartosc, err := strconv.Atoi(wartStr)
	if err != nil {
		return errors.New("strconv.Atoi(" + wartStr + "): " + err.Error())
	}

	if wartosc < 1 {
		return errors.New("wartość jest wymagana i musi być większa od zera: " + strconv.Itoa(wartosc))
	}
	return nil
}

func indexKolPoNazwie(wiersz []string, nrWiersza int, nazwaKol string) (int, error) {
	idxKol := -1
	for idx := 0; idx < len(wiersz); idx++ {
		czyNazwaKol := strings.Contains(strings.ToLower(wiersz[idx]), nazwaKol)
		if nrWiersza == 0 && czyNazwaKol {
			idxKol = idx
			break
		}
	}
	if idxKol == -1 {
		nrWierszaStr := strconv.Itoa(nrWiersza)
		return idxKol, errors.New("indexKolPoNazwie(): brak indeksu dla kolumny: " + nazwaKol + "; nrWiersza:" + nrWierszaStr)
	}
	return idxKol, nil
}

func usunPustyWiersz(p *excelize.File, sh string, wiersze [][]string, nazwa string) error {

	liczbaKolumn := len(wiersze[0])

	for nrWiersza, wiersz := range wiersze {
		if czyPystyWiersz(wiersz, liczbaKolumn) {
			log.Println("Wiersz " + strconv.Itoa(nrWiersza+1) + " jest pusty i będzie usuwany!!!")
			p.RemoveRow(sh, nrWiersza)

			err := zapiszExcela(p, nazwa)
			if err != nil {
				return err
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

func zapiszExcela(p *excelize.File, nazwa string) error {
	err := p.SaveAs(nazwa)
	if err != nil {
		return errors.New("xslx.SaveAs(): " + err.Error())
	}
	return nil
}
