package main

import (
	"fmt"
	"time"
)

func main() {
	go powiedz("Witaj")
	powiedz("  Świecie")
}

func powiedz(slowo string) {
	for i := 0; i < 5; i++ {
		time.Sleep((100 * time.Millisecond))
		fmt.Println(slowo)
	}
}
