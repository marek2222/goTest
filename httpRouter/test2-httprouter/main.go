package main

import (
	"log"
	"net/http"

	// "git.r.pl/_GoMoje/test2-httprouter/simpleUsage"

	"git.r.pl/_GoMoje/test2-httprouter/basicAuthentication"
)

func main() {
	// router := simpleUsage.SimpleUsage()
	router := basicAuthentication.BasicAuthentication()

	// user := "gordon"
	// pass := "secret!"

	// router := httprouter.New()
	// router.GET("/", Index)
	// router.GET("/protected/", BasicAuth(Protected, user, pass))

	log.Fatal(http.ListenAndServe(":8080", router))
}
