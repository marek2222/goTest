// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// Echo3 wyświetla swoje argumenty wiersza poleceń.
package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

//!+
// func main() {
// 	fmt.Println(strings.Join(os.Args[1:], " "))
// }

// pokazuje nazwę wywołującego go polecenia
func main() {
	fmt.Println(strings.Join(os.Args[0:], "; "))

	fmt.Println("\nIndex: wartość")
	for idx, arg := range os.Args {
		fmt.Printf("    %d: %s/n", idx, arg)
	}
	fmt.Println()

	start := time.Now()
	fmt.Println(strings.Join(os.Args[0:], "; "))
	fmt.Printf("Upłyneło:  %.2f sek \n", time.Since(start).Seconds())
	fmt.Printf("Upłyneło:  %d ms\n", time.Since(start).Milliseconds())

}
