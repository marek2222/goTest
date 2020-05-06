package main

import (
	"log"
	"net/http"

	"github.com/goTest/IO/CSVFile/servingAFile/servingAFile"
)

func main() {
	// 3: simpleUsage
	router := servingAFile.ServingAFile()

	log.Fatal(http.ListenAndServe(":8080", router))
}
