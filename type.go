package go_scrape_dmm_com

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
	ActorList            []*Actor
	SampleImageList      []*SampleImage
}

type Actor struct {
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

	result.ItemCode = getItemCode(url)
	result.Title = getTitle(doc)
	result.PackageImageThumbURL = getPackageImageThumbURL(doc, result.ItemCode)
	result.PackageImageURL = getPackageImageURL(doc, result.ItemCode)
	result.ActorList = getActorList(doc)
	result.SampleImageList = getSampleImageList(doc)

	return &result
}

func getItemCode(url string) string {
	cidMatcher := regexp.MustCompile(`cid=([^/]+)`)
	itemCode := cidMatcher.FindString(url)
	itemCode = cidMatcher.ReplaceAllString(itemCode, "$1")
	return itemCode
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
		if(exists) {
			packageImageThumbURL = imgSrc
		}
	})
	return packageImageThumbURL
}

func getPackageImageURL(doc *goquery.Document, itemCode string) string {
	packageImageURL := ""
	doc.Find("#" + itemCode).Each(func(index int, selection *goquery.Selection) {
		aHref, exists := selection.Attr("href")
		if(exists) {
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
		if(exists) {
			actor.ListPageURL = baseDomain + href
		}

		actorList = append(actorList, &actor)
	})

	return actorList
}

func getSampleImageList(doc *goquery.Document) []*SampleImage {
	var sampleImageList []*SampleImage

	sampleImageURLMatcher := regexp.MustCompile(`([^-]+)(-\d+\..+)`)

	doc.Find("#sample-image-block > a").Each(func(index int, selection *goquery.Selection) {
		sampleImage := SampleImage{}

		imgSrc, exists := selection.Find("img").First().Attr("src")
		if(exists) {
			sampleImage.ImageThumbURL = imgSrc
			
			imageURL :=
				sampleImageURLMatcher.ReplaceAllString(imgSrc, "$1") + "jp" +
				sampleImageURLMatcher.ReplaceAllString(imgSrc, "$2")
		
			sampleImage.ImageURL = imageURL
		}

		sampleImageList = append(sampleImageList, &sampleImage)
	})

	return sampleImageList
}