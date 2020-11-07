package main

import (
	"errors"
	"log"
	"strconv"
	"strings"

	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/marek2222/goTest/IO/excel/excelize1/hlp"
)

// SprawdzExcela sprawdza pliki excel
func SprawdzExcela(rozszerz, katalog string) error {

	pliki, err := hlp.ListOfFilesSortByName(rozszerz, katalog)
	if err != nil {
		return err
	}

	for _, plik := range pliki {
		err := sprawdzDanyPlik(plik)
		if err != nil {
			return errors.New("Błąd '" + plik + "': " + err.Error())
		}
	}
	return nil
}

func sprawdzDanyPlik(plikPelny string) error {

	p, err := excelize.OpenFile(plikPelny)
	if err != nil {
		return errors.New("Błąd przy otwieraniu pliku'" + plikPelny + "': " + err.Error())
	}

	log.Println("Skoroszyty w pliku:")

	for index, sh := range p.GetSheetMap() {
		wiersze := p.GetRows(sh)

		liczbaKolumn := len(wiersze[0])
		liczbaWierszy := len(wiersze)
		log.Println("Arkusz: index:", index, ";  nazwa:", sh)
		log.Println("   Liczba kolumn i wierszy:", liczbaKolumn, liczbaWierszy)
		log.Println("")

		for iWie, wiersz := range wiersze {
			if iWie < 2 {

				//sprawdzenie
				if czyPystyWiersz(wiersz, liczbaKolumn) {
					log.Println("Uwaga!!! Wiersz " + strconv.Itoa(iWie+1) + " jest pusty...")
					//sh.RemoveRowAtIndex(i)
				}
			}
		}
		log.Println("--------------------------------")
	}
	// Save xlsx file by the given path.
	if err := p.SaveAs(plikPelny); err != nil {
		return errors.New("Błąd przy zapisie pliku'" + plikPelny + "': " + err.Error())
	}

	return nil
}

// sprawdza czy jest pusty wiersz
func czyPystyWiersz(wiersz []string, liczbaKolumn int) bool {
	czyPystyWiersz := true
	var lista []string
	kol := liczbaKolumn
	if kol > 30 {
		kol = 30
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

// func sprawdzDanyPlik(plikPelny string) error {
//  // index := strings.LastIndex(plikPelny, "\\") + 1
//  // plik := plikPelny[index:] + "__"

//  wb, err := xlsx.OpenFile(plikPelny)
//  if err != nil {
//      return err
//  }

//  log.Println("Skoroszyty w pliku:")
//  for _, sh := range wb.Sheets {
//      liczbaKolumn := sh.MaxCol
//      liczbaWierszy := sh.MaxRow
//      log.Println("Nazwa arkusza:", sh.Name)
//      log.Println("Liczba kolumn i wierszy:", liczbaKolumn, liczbaWierszy)
//      log.Println("")

//      for i := 0; i < 2; i++ {
//          wiersz, err := sh.Row(i)
//          if err != nil {
//              return err
//          }

//          //sprawdzenie
//          czyPusty, err := czyPystyWiersz(wiersz, liczbaKolumn)
//          if err != nil {
//              return err
//          }
//          if czyPusty {
//              log.Println("Uwaga!!! Wiersz " + strconv.Itoa(i) + " jest pusty...")
//              sh.RemoveRowAtIndex(i)
//              //sh.Close()
//              plik := strings.ReplaceAll(plikPelny, "zwroty62.xlsx", "zwroty62aa.xlsx")
//              err = wb.Save(plik)
//              if err != nil {
//                  return errors.New("Błąd wb.Save(" + plikPelny + "): " + err.Error())
//              }

//          }
//          //log.Println(wiersz)

//          wiersz.ForEachCell(cellVisitor)
//          log.Println("--------------------")
//      }

//      // sh.ForEachRow(rowVisitor)
//  }
//  log.Println("----")

//  return nil
// }

// // sprawdza czy jest pusty wiersz
// func czyPystyWiersz(wiersz *xlsx.Row, liczbaKolumn int) (bool, error) {
//  czyPystyWiersz := true
//  var lista []string
//  kol := liczbaKolumn
//  if kol > 30 {
//      kol = 30
//  }
//  for i := 0; i < kol; i++ {
//      komorka, err := wiersz.GetCell(i).FormattedValue()
//      if err != nil {
//          return czyPystyWiersz, errors.New("Błąd czyPystyWiersz(): " + err.Error())
//      }
//      lista = append(lista, komorka)
//  }
//  wartosc := strings.Join(lista, "")
//  if len(wartosc) > 0 {
//      czyPystyWiersz = false
//  }

//  return czyPystyWiersz, nil
// }

// func cellVisitor(c *xlsx.Cell) error {
//  value, err := c.FormattedValue()
//  if err != nil {
//      log.Println(err.Error())
//  } else {
//      log.Println("Cell value:", value)
//  }
//  return err
// }

// func cellValue(c *xlsx.Cell) (string, error) {
//  value, err := c.FormattedValue()
//  if err != nil {
//      return "", err
//  }
//  return value, nil
// }

// func rowVisitor(r *xlsx.Row) error {
//  return r.ForEachCell(cellVisitor)
// }

// func rowStuff() {
//     filename := "samplefile.xlsx"
//     wb, err := xlsx.OpenFile(filename)
//     if err != nil {
//         panic(err)
//     }
//     sh, ok := wb.Sheet["Sample"]
//     if !ok {
//         panic(errors.New("Sheet not found"))
//     }
//     log.Println("Max row is", sh.MaxRow)
// }
