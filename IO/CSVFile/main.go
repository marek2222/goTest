package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

var data = [][]string{{"Line1", "Hello Readers of"}, {"Line2", "golangcode.com"}, {"Line3", "golangcode3.com"}}

func main() {
	file, err := os.Create("result.csv")
	checkError("Cannot create file", err)
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, value := range data {
		fmt.Printf("len=%d cap=%d %v\n", len(value), cap(value), value)
		//value = append(value, )
		err := writer.Write(value)
		checkError("Cannot write the file", err)
	}
}

func checkError(message string, err error) {
	if err != nil {
		log.Fatal(message, err)
	}
}
