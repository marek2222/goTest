package pakiet

import "fmt"

// Pakiet opis
func Pakiet() string {
	ret := ""
	for i := 1; i < 10; i++ {
		ret += fmt.Sprintf("%d ", i)
	}
	return ret
}
