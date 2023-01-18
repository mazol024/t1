package main

import (
	"fmt"
	"strings"

	"github.com/gocolly/colly"
)

func getThemes() map[string]string {
	myarrya := map[string]string{}
	// myarrya := []string{}
	c := colly.NewCollector(colly.AllowedDomains("fishki.net"))

	c.OnHTML(".content__title",
		func(e *colly.HTMLElement) {
			title := e.ChildText("a")
			themhr := e.ChildAttr("a", "href")
			if strings.Contains(themhr, "html") {
				// myarrya = append(myarrya, themhr)
				myarrya[title] = themhr
				fmt.Println(title)
				fmt.Println(themhr)
			}
		})
	c.Visit("https://fishki.net")
	return myarrya
}
func runHttp3(visit string) []string {
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

	c.OnHTML(".paragraph",
		func(e *colly.HTMLElement) {
			prgph := e.Text
			fmt.Println(prgph)
			fmt.Println("+++++++++")
		})

	c.OnHTML(".picture-relative",
		func(e *colly.HTMLElement) {
			adrw := e.ChildAttr("a", "href")
			themw := e.ChildAttr("img", "alt")
			fmt.Println(adrw)
			fmt.Println(themw)
			if strings.Contains(adrw, ".jpg") || strings.Contains(adrw, ".png") {
				myarray = append(myarray, adrw)
			} else if strings.Contains(adrw, ".html") {
				myarray = append(myarray, "https://fishki.net"+adrw)

			}
		})

	c.Visit("https://fishki.net" + visit)
	// c.Visit("https://fishki.net/4340725-zachem-nuzhny-byli-chepchiki-dlja-sna.html")

	return myarray
}
