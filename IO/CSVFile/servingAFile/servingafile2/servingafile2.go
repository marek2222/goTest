package servingafile2

import (
	"fmt"
	"net/http"
	"path"

	"github.com/julienschmidt/httprouter"
)

// ServingAFile call httprouter
func ServingAFile() *httprouter.Router {
	router := httprouter.New()
	router.GET("/", Index)
	return router
}

// Index function
func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// Assuming you want to serve a photo at 'images/foo.png'
	fp := path.Join("./images", "foo.jpg")
	odp := fmt.Sprintf("%s", fp)
	w.Write([]byte(odp))
	http.ServeFile(w, r, fp)
}
