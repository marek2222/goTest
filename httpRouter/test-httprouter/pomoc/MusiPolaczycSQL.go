package pomoc

import (
	"log"
	"strings"

	"github.com/jmoiron/sqlx"
)

// MusiPolaczycSQL tworzy obiekt db
func MusiPolaczycSQL(connStr string, maxSQLConns int) *sqlx.DB {
	db, err := sqlx.Connect("mssql", connStr)
	if err != nil {
		log.Fatalf("błąd połączenia z mssql \"%s\"\n", strings.Split(connStr, "pass")[0])
	}
	db.SetMaxOpenConns(maxSQLConns)
	return db
}
