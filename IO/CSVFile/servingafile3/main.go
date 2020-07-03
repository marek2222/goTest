package main

import (
	"log"
	"net/http"
)

// TmpDir temporary directory location
const TmpDir = "github.com/marek2222/goTest/IO/CSVFile/servingAFile/servingafile3"

func main() {

	// return a `.pdf` file for `/pdf` route
	http.HandleFunc("/pdf", func(res http.ResponseWriter, req *http.Request) {
		http.ServeFile(res, req, TmpDir+"/test.pdf")
	})

	// start HTTP server with `http.DefaultServeMux` handler
	log.Fatal(http.ListenAndServe(":9000", nil))

}
