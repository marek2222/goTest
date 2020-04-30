package test

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func zapisHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	samolotID, err := strconv.Atoi(ps.ByName("samolotid"))
	if err != nil {
		log.Println("Błąd samolotID:", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	bodyJSON, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("Błąd bodyJSON:", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
	}

	err = zapiszDoBazy(samolotID, string(bodyJSON))
	if err != nil {
		log.Println(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "samolotID = %d zapisany do bazy", samolotID)
}

func zapiszDoBazy(samolotID int, bodyJSON string) error {
	updateQuery := "update miejsca.Samoloty set MiejscaJSON=? where SamolotID=?"

	_, err := db.Exec(updateQuery, bodyJSON, samolotID)
	if err != nil {
		return err
	}
	return nil
}
