package test

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func usuwanieHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	samolotID, err := strconv.Atoi(ps.ByName("samolotid"))
	if err != nil {
		log.Println("Błąd samolotID:", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	err = usunZBazy(samolotID)
	if err != nil {
		log.Println("Błąd usuwania:", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "usunąłem samolotID = %d", samolotID)
}

func usunZBazy(samolotID int) error {
	usunQuery := "update miejsca.Samoloty set MiejscaJSON=null where SamolotID=?"
	_, err := db.Exec(usunQuery, samolotID)
	if err != nil {
		return err
	}
	return nil
}
