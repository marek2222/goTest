package basicAuthentication

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// BasicAuthentication for handles:
func BasicAuthentication() *httprouter.Router {
	user := "gordon"
	pass := "secret!"

	router := httprouter.New()
	router.GET("/", Index)
	router.GET("/protected/", BasicAuth(Protected, user, pass))
	return router
}

// BasicAuth ...
func BasicAuth(h httprouter.Handle, requiredUser, requiredPassword string) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		// Get the Basic Authentication credentials
		user, password, hasAuth := r.BasicAuth()

		if hasAuth && user == requiredUser && password == requiredPassword {
			// Delegate request to the given handle
			h(w, r, ps)
		} else {
			// Request Basic Authentication otherwise
			w.Header().Set("WWW-Authenticate", "Basic realm=Restricted")
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		}
	}
}

// Index Not protected function
func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Not protected!\n")
}

// Protected Protected function
func Protected(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Protected!\n")
}
