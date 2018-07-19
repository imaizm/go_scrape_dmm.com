package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
)

func main() {
	url := "http://www.dmm.com/digital/idol/-/detail/=/cid=5013tsds42319/"

	doc, err := goquery.NewDocument(url)
	if err != nil {
		panic(err)
	}

	selection := doc.Find("#title")
	text := selection.First().Text()

	fmt.Println(text)
}