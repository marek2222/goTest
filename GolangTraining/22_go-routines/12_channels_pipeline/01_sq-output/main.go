package main

import "fmt"

func main() {
	// Set up the pipeline.
	out := gen(2, 3)
	// c := gen(2, 3)
	// out := sq(c)

	// Consume the output.
	fmt.Println(<-out) // 4
	fmt.Println(<-out) // 9
}

func gen(nums ...int) chan int {
	outCh := make(chan int)
	go func() {
		for _, n := range nums {
			outCh <- n
		}
		close(outCh)
	}()

	out := make(chan int)
	go func() {
		for n := range outCh {
			out <- n * n
		}
		close(out)
	}()
	return out
}

// func gen(nums ...int) chan int {
// 	out := make(chan int)
// 	go func() {
// 		for _, n := range nums {
// 			out <- n
// 		}
// 		close(out)
// 	}()
// 	return out
// }

// func sq(in chan int) chan int {
// 	out := make(chan int)
// 	go func() {
// 		for n := range in {
// 			out <- n * n
// 		}
// 		close(out)
// 	}()
// 	return out
// }
