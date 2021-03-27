// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

//!+

// Dup2 wyświetla liczbę wystąpień i tekst linii, które w danych wejściowych pojawiają się więcej niż raz.
// Odczytuje ze standardowego wejścia lub z listy nazwanych plików.

// go run main.go main1.txt main2.txt
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]
	fmt.Printf("Liczba: %d\n", len(files))
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			fmt.Printf("File: %s\n", arg)
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("Files: %d\t Line: %s\n", n, line)
		}
	}
}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
	// UWAGA: ignorowanie potencjalnych błędów z fukcji input.Err()
}
