package main

import (
	"bytes"
	"fmt"
	"regexp"
)

func main() {
	//stringRegularExpression()
	floatRegularExpression()

}

func stringRegularExpression() {
	match, _ := regexp.MatchString("p([a-z]+)ch", "peach")
	fmt.Println(match)

	r, _ := regexp.Compile("p([a-z]+)ch")

	fmt.Println("peach: ", r.MatchString("peach"))

	fmt.Println("FindString: ", r.FindString("peach punch"))

	fmt.Println(r.FindStringIndex("peach punch"))

	fmt.Println(r.FindStringSubmatch("peach punch"))

	fmt.Println(r.FindStringSubmatchIndex("peach punch"))

	fmt.Println(r.FindAllString("peach punch pinch", -1))

	fmt.Println(r.FindAllStringSubmatchIndex("peach punch pinch", -1))

	fmt.Println(r.FindAllString("peach punch pinch", 2))

	fmt.Println(r.Match([]byte("peach")))

	r = regexp.MustCompile("p([a-z]+)ch")
	fmt.Println(r)

	fmt.Println(r.ReplaceAllString("a peach", "<fruit>"))

	in := []byte("a peach")
	out := r.ReplaceAllFunc(in, bytes.ToUpper)
	fmt.Println(string(out))
}

func floatRegularExpression() {
	// match, _ := regexp.MatchString("([a-z]+) [+-]?([0-9]*[.])?[0-9]+  ZŁ", "WYBIERAM ZA 950 ZŁ")
	// fmt.Println(match)

	match, _ := regexp.MatchString("([A-Z][^A-Z]*)[ ][+-]?([0-9]*[.])?[0-9]+[ ]ZŁ", "WYBIERAM ZA 950 ZŁ")
	fmt.Println(match)

	r, _ := regexp.Compile("[+-]?([0-9]*[.])?[0-9]+")
	fmt.Println("MatchString: ", r.MatchString("950"))

}
