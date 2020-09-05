// https://www.it-swarm.dev/pl/go/jak-zatrzymac-goroutine/940528838/

package main

import (
	"fmt"
	"time"
)

var licznik int

func doStuff() int {
	licznik++
	return licznik
}

func main() {

	ch := make(chan int, 100)
	done := make(chan struct{})
	go func() {
		for {
			select {
			case ch <- doStuff():
			case <-done:
				close(ch)
				return
			}
			time.Sleep(100 * time.Millisecond)
		}
	}()

	go func() {
		time.Sleep(2 * time.Second)
		done <- struct{}{}
	}()

	for i := range ch {
		fmt.Println("receive value: ", i)
	}

	fmt.Println("finish: licznik:", licznik)
}
