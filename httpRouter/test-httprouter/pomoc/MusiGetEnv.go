package pomoc

import (
	"log"
	"os"
)

// MusiGetEnv pobiera zadaną env
func MusiGetEnv(env string) string {
	wartosc := os.Getenv(env)
	if wartosc == "" {
		log.Fatalf("błąd: nie ustawiono zmiennej środowiskowej \"%s\"\n", env)
	}
	return wartosc
}
