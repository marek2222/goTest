package main

import (
	"log"
	"net/http"

	"github.com/marek2222/goTest/IO/CSVFile/servingAFile/servingafile2"
)

func main() {
	// 3: simpleUsage
	router := servingafile2.ServingAFile()

	log.Fatal(http.ListenAndServe(":8080", router))
}
