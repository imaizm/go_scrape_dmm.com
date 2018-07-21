package main

import (
	"fmt"
	"strconv"
)

func main() {
	url := "http://www.dmm.com/digital/idol/-/detail/=/cid=5013tsds42319/"

	result := New(url)

	fmt.Println(result.ItemCode)
	fmt.Println(result.Title)
	fmt.Println(result.PackageImageThumbURL)
	fmt.Println(result.PackageImageURL)
	for index, value := range result.ActressList {
		fmt.Println(strconv.Itoa(index) + " : " + value.Name + " : " + value.ListPageURL)
	}
	for index, value := range result.SampleImageList {
		fmt.Println(strconv.Itoa(index) + " : " + value.ImageThumbURL + " : " + value.ImageURL)
	}
}