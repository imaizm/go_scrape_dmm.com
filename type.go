package main

import (
	"regexp"
	"github.com/PuerkitoBio/goquery"
)

const baseDomain = "http://www.dmm.com"

type ItemOfDmmComIdol struct {
	ItemCode             string
	Title                string
	PackageImageThumbURL string
	PackageImageURL      string
	ActressList          []*Actress
	SampleImageList      []*SampleImage
}

type Actress struct {
	ListPageURL string
	Name string
}

type SampleImage struct {
	ImageThumbURL string
	ImageURL      string
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
			actress.ListPageURL = baseDomain + href
		}

		result.ActressList = append(result.ActressList, &actress)
	})


	sampleImageUrlMatcher := regexp.MustCompile(`([^-]+)(-\d+\..+)`)

	doc.Find("#sample-image-block > a").Each(func(index int, selection *goquery.Selection) {
		sampleImage := SampleImage{}

		src, exists := selection.Find("img").First().Attr("src")
		if(exists) {
			sampleImage.ImageThumbURL = src
			
			imageURL := sampleImageUrlMatcher.ReplaceAllString(src, "$1") + "jp" + sampleImageUrlMatcher.ReplaceAllString(src, "$2")
		
			sampleImage.ImageURL = imageURL
		}

		result.SampleImageList = append(result.SampleImageList, &sampleImage)
	})

	return &result
}