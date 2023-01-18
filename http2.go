package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

func main() {

	c := colly.NewCollector(
		colly.AllowedDomains("fishki.net"),
	)
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})
	c.OnResponse(func(r *colly.Response) {
		fmt.Println(r.StatusCode)
	})
	c.OnHTML("title", func(e *colly.HTMLElement) {
		fmt.Println(e.Text)
	})
	c.OnHTML("post-pic-inline", func(e *colly.HTMLElement) {
		a1 := e.ChildAttr("src", "text")
		a2 := e.ChildAttr("alt", "href")
		fmt.Println(a1, a2)
	})

	c.Visit("https://fishki.net/")
}
