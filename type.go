package goScrapeDmmCom

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/imaizm/go_scrape_dmm-common"
)

const baseDomain = "http://www.dmm.com"

// ItemOfDmmComIdol : ItemOfDmmComIdol Info Struct
type ItemOfDmmComIdol struct {
	ItemCode             string
	Title                string
	PackageImageThumbURL string
	PackageImageURL      string
	ActorList            []*Actor
	SampleImageList      []*goScrapeDmmCommon.SampleImage
}

// Actor : Actor Info Struct
type Actor struct {
	ListPageURL string
	Name        string
}

// New : create ItemOfDmmComIdol struct from url
func New(url string) *ItemOfDmmComIdol {

	doc, err := goquery.NewDocument(url)
	if err != nil {
		panic(err)
	}

	result := ItemOfDmmComIdol{}

	result.ItemCode = getItemCode(url)
	result.Title = getTitle(doc)
	result.PackageImageThumbURL = getPackageImageThumbURL(doc, result.ItemCode)
	result.PackageImageURL = getPackageImageURL(doc, result.ItemCode)
	result.ActorList = getActorList(doc)
	result.SampleImageList = getSampleImageList(doc)

	return &result
}

func getItemCode(url string) string {
	return goScrapeDmmCommon.GetItemCodeFromURL(url)
}

func getTitle(doc *goquery.Document) string {
	selection := doc.Find("#title")
	title := selection.First().Text()
	return title
}

func getPackageImageThumbURL(doc *goquery.Document, itemCode string) string {
	packageImageThumbURL := ""
	doc.Find("#package-src-" + itemCode).Each(func(index int, selection *goquery.Selection) {
		imgSrc, exists := selection.Attr("src")
		if exists {
			packageImageThumbURL = imgSrc
		}
	})
	return packageImageThumbURL
}

func getPackageImageURL(doc *goquery.Document, itemCode string) string {
	packageImageURL := ""
	doc.Find("#" + itemCode).Each(func(index int, selection *goquery.Selection) {
		aHref, exists := selection.Attr("href")
		if exists {
			packageImageURL = aHref
		}
	})
	return packageImageURL
}

func getActorList(doc *goquery.Document) []*Actor {
	var actorList []*Actor

	doc.Find("table.mg-b20").First().Find("a[href *= 'article=actor']").Each(func(index int, selection *goquery.Selection) {
		actor := Actor{}
		actor.Name = selection.Text()

		href, exists := selection.Attr("href")
		if exists {
			actor.ListPageURL = baseDomain + href
		}

		actorList = append(actorList, &actor)
	})

	return actorList
}

func getSampleImageList(doc *goquery.Document) []*goScrapeDmmCommon.SampleImage {
	return goScrapeDmmCommon.GetSampleImageList(doc)
}
