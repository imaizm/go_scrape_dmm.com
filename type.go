package main

import (
	"github.com/PuerkitoBio/goquery"
)

type ItemOfDmmComIdol struct {
	ItemCode string
	Title string
}


func NewHoge() *ItemOfDmmComIdol {
	url := "http://www.dmm.com/digital/idol/-/detail/=/cid=5013tsds42319/"

	doc, err := goquery.NewDocument(url)
	if err != nil {
		panic(err)
	}

	result := ItemOfDmmComIdol{}
	selection := doc.Find("#title")
	result.Title = selection.First().Text()

	return &result
}