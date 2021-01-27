package main

import (
	"fmt"
	"time"
)

type osoba struct {
	nazwa *nazwa
	czas  *time.Time
}

type nazwa struct {
	imie string
}

func main() {
	k := osoba{
		&nazwa{"krzyś"},
		addr(time.Now()),
	}
	fmt.Println(k.nazwa.imie, k.czas)
}

func addr(t time.Time) *time.Time {
	return &t
}
