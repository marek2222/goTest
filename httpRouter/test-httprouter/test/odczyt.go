package test

import (
	"log"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func odczytHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	samolotID, err := strconv.Atoi(ps.ByName("samolotid"))
	if err != nil {
		log.Println("Błąd samolotID:", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	bodyJSON, err := wczytajZBazy(samolotID)
	if err != nil {
		log.Println("Błąd wczytywania:", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(bodyJSON)
}

func wczytajZBazy(samolotID int) ([]byte, error) {
	selectQuery := "select s.MiejscaJSON from miejsca.Samoloty s where s.SamolotID=?"
	var miejscaJSON *string
	err := db.QueryRow(selectQuery, samolotID).Scan(&miejscaJSON)
	if err != nil {
		return nil, err
	}
	if miejscaJSON == nil {
		nullJSON := "{}"
		miejscaJSON = &nullJSON
	}
	return []byte(*miejscaJSON), nil
}
