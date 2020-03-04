package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"git.r.pl/rpl-api/go-moje/pakiet"
)

func main() {
	port := os.Getenv("PORT")

	fmt.Println("Dzień dobry")

	http.HandleFunc("/nasza", naszHandler)
	fmt.Println("Startuję serwer http na porcie", port)

	http.ListenAndServe(port, nil)
}

func naszHandler(w http.ResponseWriter, r *http.Request) {
	cos := pakiet.Pakiet()
	log.Println(cos)
	fmt.Fprintln(w, "Dzień dobry.")
	fmt.Fprintln(w, cos)
	jsn, _ := json.Marshal(cos)
	w.Write(jsn)

}
