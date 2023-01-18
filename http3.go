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

	c.OnHTML(".content__title",
		func(e *colly.HTMLElement) {
			// price := e.Text
			price := e.ChildText("a span")
			title := e.ChildAttr("a", "href")
			fmt.Println(price)
			fmt.Println(title)
			fmt.Println("+++++++++")

		})
	c.OnHTML(".picture-relative",
		func(e *colly.HTMLElement) {
			// price := e.Text
			// price := e.ChildText("a")
			price1 := e.ChildAttr("img", "src")
			price2 := e.ChildAttr("img", "alt")
			// price := e.Attr("src")

			// fmt.Println(price)
			fmt.Println(price1)
			fmt.Println(price2)
		})

	// c.Visit("https://fishki.net/4340457-foto-zabroshennogo-boinga-kotoryj-prevratili-v-krutuju-villu.html")
	c.Visit("https://fishki.net/auto/4340490-podborka-jumora-dlja-avtoljubitelej.html")
	// printimg(images1, site)
	// drawinghttp(images1, site)
}
