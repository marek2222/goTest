package main

import (
	"fmt"

	winniepooh "github.com/goTest/GolangTraining/02_package/icomefromalaska"
	"github.com/goTest/GolangTraining/02_package/stringutil"
	//someAlias "github.com/GoesToEleven/GolangTraining/02_package/icomefromalaska"
)

func main() {
	fmt.Println(stringutil.Reverse("!oG ,olleH"))
	fmt.Println(stringutil.MyName)
	fmt.Printf("BearName: %s", winniepooh.BearName)
}
