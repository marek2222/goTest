// https://www.it-swarm.dev/pl/go/jak-zatrzymac-goroutine/940528838/

package main

import (
	"context"
	"fmt"
	"time"
)

var counter int

func show() {
	counter++
	fmt.Println("counter", counter)
}

func main() {
	forever := make(chan struct{})
	ctx, cancel := context.WithCancel(context.Background())

	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done(): // if cancel() execute
				forever <- struct{}{}
				return
			default:
				show()
			}

			time.Sleep(500 * time.Millisecond)
		}
	}(ctx)

	go func() {
		time.Sleep(3 * time.Second)
		cancel()
	}()

	<-forever
	fmt.Println("finish: counter", counter)
}
