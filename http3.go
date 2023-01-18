package main

import (
	"fmt"
	"strings"

	"github.com/gocolly/colly"
)

func runHttp3() []string {
	myarray := []string{}
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

	c.OnHTML(".picture-relative",
		func(e *colly.HTMLElement) {
			// title := e.ChildText(".paragraph  ")
			title := e.ChildAttr("a", "href")
			// price := e.Text
			// price := e.ChildAttr("img", "href")
			// price := e.ChildText("img")
			fmt.Println(title)
			// fmt.Println(price)
			if strings.Contains(title, ".jpg") || strings.Contains(title, ".png") {
				myarray = append(myarray, title)
			}
		})

	c.Visit("https://fishki.net/4340823-i-smeh-i-greh-ili-tipichnye-rabochie-budni-posle-dlinnyh-vyhodnyh.html")
	// printimg(images1, site)
	// drawinghttp(images1, site)
	return myarray
}
