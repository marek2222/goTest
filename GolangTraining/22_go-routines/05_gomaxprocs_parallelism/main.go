package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var wg sync.WaitGroup

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	fmt.Printf("Współbieżność rozłożono na %d procesorów...", runtime.NumCPU())
	fmt.Println()
	fmt.Println("Funckja foo() śpi przez 3, a bar() przez 20 milisekund")
	fmt.Println("")
	fmt.Println("Początek współbieżności")
	wg.Add(2)
	go foo()
	go bar()
	wg.Wait()
	fmt.Println("Koniec współbieżności")
}

func foo() {
	for i := 0; i < 17; i++ {
		fmt.Println("Foo:", i)
		time.Sleep(3 * time.Millisecond)
	}
	wg.Done()
}

func bar() {
	for i := 0; i < 17; i++ {
		fmt.Println("Bar:", i)
		time.Sleep(20 * time.Millisecond)
	}
	wg.Done()
}
