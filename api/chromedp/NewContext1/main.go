// Command click is a chromedp example demonstrating how to use a selector to
// click on an element.
package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"

	"github.com/chromedp/chromedp"
)

func main() {
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	ts := httptest.NewServer(writeHTML(`
	<body>
	<p id="content" onclick="changeText()">Original content.</p>
	<script>
	function changeText() {
		document.getElementById("content").textContent = "New content!"
	}
	</script>
	</body>
		`))
	defer ts.Close()

	var outerBefore, outerAfter string
	if err := chromedp.Run(ctx,
		chromedp.Navigate(ts.URL),
		chromedp.OuterHTML("#content", &outerBefore),
		chromedp.Click("#content", chromedp.ByID),
		chromedp.OuterHTML("#content", &outerAfter),
	); err != nil {
		panic(err)
	}
	fmt.Println("OuterHTML before clicking:")
	fmt.Println(outerBefore)
	fmt.Println("OuterHTML after clicking:")
	fmt.Println(outerAfter)
}

func writeHTML(content string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//utf-8
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		io.WriteString(w, strings.TrimSpace(content))
	})
}
