package main

import {
	"fmt"
	"github.com/fatih/structs"
}

func main() {
	data := []MyStruct{MyStruct{}, MyStruct{}}
	buff := &bytes.Buffer{}
	w := struct2csv.NewWriter(buff)
	err := w.WriteStructs(data)
	if err != nil {
			// handle error
	}
}
