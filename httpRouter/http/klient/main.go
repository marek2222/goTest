package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

//program właściwy
func main() {
	url, fname := setup()
	//wykonujemy zapytanie
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	//czytanie odpowiedzi
	rbody := resp.Body
	buf, rerr := ioutil.ReadAll(rbody)
	rbody.Close()

	if rerr != nil {
		fmt.Println("Cannot read response body")
		fmt.Println(rerr)
		os.Exit(1)
	}
	//zapis odpowiedzi do pliku
	ioutil.WriteFile(fname, buf, 0666)
}

//przetwarzanie argumentów i przełączników z linii poleceń
func setup() (url, fname string) {
	//oczekujemy przełącznika -url
	var urlopt = flag.String("url", "", "Resource URL")
	//oczekujemy nazwy pliku do którego zapiszemy pobrany zasób w opcji -o
	var fnameopt = flag.String("o", "out", "Output filename")
	flag.Parse() //przetwarzanie

	fname = *fnameopt
	urlopt2 := *urlopt
	var urlarg = flag.Arg(0) //url może być argumentem a nie opcją

	// if urlarg == "" || urlopt2 == "" {
	// 	flag.PrintDefaults()
	// 	return
	// }

	//url z argumentu jest ważniejszy niż z przełącznika
	if len(urlarg) > 0 {
		url = urlarg
	} else if len(urlopt2) > 0 {
		url = urlopt2
	} else {
		fmt.Println("No URL provided")
		os.Exit(1)
	}
	return //zwracamy wartości w domyślnych zmiennych
}
