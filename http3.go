package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

func main() {

	c := colly.NewCollector(colly.AllowedDomains("fishki.net"))

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Println(r.StatusCode)
	})

	c.OnHTML("title", func(e *colly.HTMLElement) {
		fmt.Println(e.Text)
	})

	c.OnHTML(".content__text  ",
		func(e *colly.HTMLElement) {
			title := e.ChildText(".paragraph  ")
			// title := e.ChildAttr("a", "href")
			// price := e.Text
			// price := e.ChildAttr("img", "href")
			// price := e.ChildText("img")
			fmt.Println(title)
			// fmt.Println(price)
		})

	c.Visit("https://fishki.net")
	// printimg(images1, site)
	// drawinghttp(images1, site)
}
