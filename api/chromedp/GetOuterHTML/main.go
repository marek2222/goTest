package main

import (
	"context"
	"fmt"

	"github.com/chromedp/cdproto/dom"
	"github.com/chromedp/chromedp"
)

func main() {
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	var res string

	err := chromedp.Run(ctx,
		chromedp.Navigate(`https://github.com/chromedp/chromedp`),
		chromedp.ActionFunc(func(ctx context.Context) error {
			node, err := dom.GetDocument().Do(ctx)
			if err != nil {
				return err
			}
			res, err = dom.GetOuterHTML().WithNodeID(node.NodeID).Do(ctx)
			return err
		}),
	)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(res)
}
