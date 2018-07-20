package main

import (
	"regexp"
	"github.com/PuerkitoBio/goquery"
)

type ItemOfDmmComIdol struct {
	ItemCode             string
	Title                string
	PackageImageThumbURL string
	PackageImageURL      string
	ActressList          []*Actress
}

type Actress struct {
	Id string
	Name string
}


func New(url string) *ItemOfDmmComIdol {

	doc, err := goquery.NewDocument(url)
	if err != nil {
		panic(err)
	}


	result := ItemOfDmmComIdol{}

	cidMatcher := regexp.MustCompile(`cid=([^/]+)`)
	itemCode := cidMatcher.FindString(url)
	itemCode = cidMatcher.ReplaceAllString(itemCode, "$1")
	result.ItemCode = itemCode

	selection := doc.Find("#title")
	result.Title = selection.First().Text()

	doc.Find("#package-src-"+itemCode).Each(func(index int, selection *goquery.Selection) {
		img_src, exists := selection.Attr("src")
		if(exists) {
			result.PackageImageThumbURL = img_src
		}
	})

	doc.Find("#"+itemCode).Each(func(index int, selection *goquery.Selection) {
		a_href, exists := selection.Attr("href")
		if(exists) {
			result.PackageImageURL = a_href
		}
	})

	doc.Find("table.mg-b20").First().Find("a[href *= 'article=actor']").Each(func(index int, selection *goquery.Selection) {
		actress := Actress{}
		actress.Name = selection.Text()

		href, exists := selection.Attr("href")
		if(exists) {
			actress.Id = href
		}

		result.ActressList = append(result.ActressList, &actress)
	})

	return &result
}