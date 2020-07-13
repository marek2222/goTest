// Command text is a chromedp example demonstrating how to extract text from a
// specific element.
package main

import (
	"context"
	"log"
	"strings"

	"github.com/chromedp/chromedp"
)

func main() {
	// create context
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	// run task list
	var res string
	err := chromedp.Run(ctx,
		chromedp.Navigate(`https://biletyczarterowe2.r.pl/destynacja?data=2020-07-17&dokad%5B%5D=RHO&idPrzylot=42816_175952&idWylot=175974&oneWay=false&pakietIdPrzylot=42816_175952&pakietIdWylot=42816_175974&przylotDo&przylotOd&wiek%5B%5D=1989-10-30&wiek%5B%5D=1989-10-30&wiek%5B%5D=2018-07-10&wiek%5B%5D=2018-07-10&wylotDo&wylotOd`),
		chromedp.Text(`#app div.destynacja section.section.accent-bg.bilety-section.bilety-section a.button.kupuje-button`, &res, chromedp.NodeVisible, chromedp.ByID),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(strings.TrimSpace(res))
}
