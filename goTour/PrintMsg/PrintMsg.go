package PrintMsg

import "fmt"

// Print show message with title
func Print(title string, message string) {
	fmt.Println("  "+title, ":    ", message)
}
