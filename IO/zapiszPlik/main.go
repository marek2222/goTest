package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	tekst := "tekst do zapisu"
	fileName := "app.html"
	err := saveToFile(tekst, fileName)
	if err != nil {
		fmt.Println("Error: ", err)
	}
}

func saveToFile(tekst, fileName string) error {
	f, err := os.Create(fileName)
	if err != nil {
		return err
	}
	// It’s idiomatic to defer a Close immediately after opening a file.
	defer f.Close()

	w := bufio.NewWriter(f)
	n4, err := w.WriteString(tekst + "\n")
	if err != nil {
		return err
	}
	fmt.Printf("Save %d bytes\n", n4)

	//Use Flush to ensure all buffered operations have been applied to the underlying writer.
	w.Flush()
	return nil
}
