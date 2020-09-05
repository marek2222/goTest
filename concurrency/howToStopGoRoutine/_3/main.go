// https://www.it-swarm.dev/pl/go/jak-zatrzymac-goroutine/940528838/

package main

import (
	"fmt"
	"time"

	"gopkg.in/tomb.v2"
)

//Proc ...
type Proc struct {
	Tomb tomb.Tomb
}

// Exec ...
func (proc *Proc) Exec() {
	defer proc.Tomb.Done() // Must call only once
	for {
		select {
		case <-proc.Tomb.Dying():
			return
		default:
			time.Sleep(300 * time.Millisecond)
			fmt.Println("Loop the loop")
		}
	}
}

func main() {
	proc := &Proc{}
	go proc.Exec()
	time.Sleep(1 * time.Second)
	proc.Tomb.Kill(fmt.Errorf("Death from above"))
	err := proc.Tomb.Wait() // Will return the error that killed the proc
	fmt.Println(err)
}
