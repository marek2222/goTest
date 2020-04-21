package main

import (
	"fmt"
	"log"
	"net/http"

	"git.r.pl/_GoMoje/test-httprouter/pomoc"
	"git.r.pl/bilety-czarterowe/tools/samoloty-miejsca/test"
	_ "github.com/denisenkom/go-mssqldb"
	"github.com/julienschmidt/httprouter"
)

const addr = ":80"

func main() {
	fmt.Println("Dzień dobry....")

	db := pomoc.MusiPolaczycSQL(pomoc.MusiGetEnv("SQL_DB"), 5)

	ruter := httprouter.New()
	ruter.GET("/", func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		fmt.Fprint(w, "Samoloty-miejsca. Cześć.")
	})
	test.Start(ruter, db)

	fmt.Printf("http.ListenAndServe \"%s\"\n", addr)
	log.Fatal(http.ListenAndServe(addr, ruter))
}
