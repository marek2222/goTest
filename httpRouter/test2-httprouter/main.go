package main

import (
	"log"
	"net/http"

	"github.com/goTest/httpRouter/test2-httprouter/servingAFile"
)

func main() {
	// 3: simpleUsage
	router := servingAFile.ServingAFile()

	// // 2: basicAuthentication
	// router := basicAuthentication.BasicAuthentication()

	// user := "gordon"
	// pass := "secret!"

	// router := httprouter.New()
	// router.GET("/", Index)
	// router.GET("/protected/", BasicAuth(Protected, user, pass))

	// // 1: simpleUsage
	// router := simpleUsage.SimpleUsage()

	log.Fatal(http.ListenAndServe(":8080", router))
}
