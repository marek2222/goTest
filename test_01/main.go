package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	port := "80" //os.Getenv("PORT")
	fmt.Println(port)

	fmt.Println("Dzień dobry")

	http.HandleFunc("/nasza", naszHandler)
	fmt.Println("Startuję serwer http na porcie", port)

	http.ListenAndServe(port, nil)
}

func naszHandler(w http.ResponseWriter, r *http.Request) {
	cos := ""
	for i := 1; i < 10; i++ {
		cos += fmt.Sprintf("%d ", i)
	}
	log.Println(cos)
	// fmt.Fprintln(w, "Dzień dobry.")
	// fmt.Fprintln(w, cos)
	// jsn, _ := json.Marshal(cos)
	// w.Write(jsn)

}

// Pakiet opis
func Pakiet() string {
	ret := ""
	for i := 1; i < 10; i++ {
		ret += fmt.Sprintf("%d ", i)
	}
	return ret
}
