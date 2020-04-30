package simpleUsage

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// SimpleUsage call httprouter
func SimpleUsage() *httprouter.Router {
	router := httprouter.New()
	router.GET("/", Index)
	router.GET("/hello/:name", Hello)
	return router
}

// Index function
func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Witaj!\n")
}

// Hello funmction
func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "Cześć, %s!\n", ps.ByName("name"))
}
