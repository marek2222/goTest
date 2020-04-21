package stringutil

import (
	"fmt"
)

func reverseTwo(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		// e1, e2 := r[i], r[j]
		// r := rune('a')
		fmt.Printf("%d; %d; %d;", i, j, len(r)-1)
		fmt.Println(r)

		// e1, e2 := r[i], r[j]
		// fmt.Printf(strconv.Itoa(e1))
		// fmt.Printf(strconv.Itoa(e2))
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

// this demonstrates how an unexported function
// can be used by an exported function in the same package
