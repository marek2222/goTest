package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("Początek współbieżności")
	wg.Add(2)
	go foo()
	go bar()
	wg.Wait()
	fmt.Println("Koniec współbieżności")
}

func foo() {
	for i := 0; i < 15; i++ {
		fmt.Println("Foo:", i)
		time.Sleep(3 * time.Millisecond)
	}
	wg.Done()
}

func bar() {
	for i := 0; i < 15; i++ {
		fmt.Println("Bar:", i)
		time.Sleep(20 * time.Millisecond)
	}
	wg.Done()
}
