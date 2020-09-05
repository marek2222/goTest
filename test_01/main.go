package main

import (
	"fmt"
	"io"
	"net/http"

	"github.com/marek2222/goTest/test_01/pakiet"
)

func main() {
	port := ":9000" //os.Getenv("PORT")
	fmt.Println(port)

	fmt.Println("Dzień dobry")

	http.HandleFunc("/nasza", naszHandler)
	fmt.Println("Startuję serwer http na porcie", port)

	http.ListenAndServe(port, nil)
}

func naszHandler(w http.ResponseWriter, r *http.Request) {
	cos := pakiet.Pakiet()
	// fmt.Println(cos)
	io.WriteString(w, cos)
	// cos := ""
	// for i := 1; i < 10; i++ {
	// 	cos += fmt.Sprintf("%d ", i)
	// }
	// log.Println(cos)

	// fmt.Fprintln(w, "Dzień dobry.")
	// fmt.Fprintln(w, cos)
	// jsn, _ := json.Marshal(cos)
	// w.Write(jsn)

}

// // Pakiet opis
// func Pakiet() string {
// 	ret := ""
// 	for i := 1; i < 10; i++ {
// 		ret += fmt.Sprintf("%d ", i)
// 	}
// 	return ret
// }
