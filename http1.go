package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func main() {
	var images []string
	// var images []fyne.URI
	urls := []string{
		"https://www.eatliver.com/",
		"https://fishki.net",
		"https://thechive.com/category/humor/funny-pictures/",
		"https://www.freepik.com",
		"https://lenta.ru",
		"https://rbc.ru",
	}
	site := urls[0]
	type c http.Client
	resp, _ := http.Get(site)
	var rows []string

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	// ...
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Status received for %s: %d\n",
		resp.Request.URL, resp.StatusCode)

	page1 := string(body)

	rows = backsearch(page1)
	fmt.Println("rows befor check", len(rows))
	for _, i := range rows {
		i = i[strings.Index(i, "http"):]
		images = append(images, i)
	}
	fmt.Println("rows after check", len(images))
	var images1 []string
	// var images1 []fyne.URI
	for _, i := range images {
		// fmt.Println(i.String())
		if notin(i, images1) {
			images1 = append(images1, i)
		}
	}
	fmt.Println(" Images2  :", len(images1))
	// printimg(images1, site)
	drawinghttp(images1, site)
}
func backsearch(s string) []string {
	var rows []string
	i := len(s)
	for i > 0 {

		endp := strings.Index(s, ".jpg")
		if endp < 0 {
			break
		}
		startp := strings.LastIndex(s[:endp], "https:")
		if startp < 0 || startp > endp {
			break
		}
		ss := s[startp : endp+4]
		rows = append(rows, ss)
		s = s[endp+4:]
		i = len(s)
	}
	return rows
}
func notin(i string, images []string) bool {
	// func notin(i fyne.URI, images []fyne.URI) bool {
	for _, p := range images {
		if p == i {
			fmt.Println(" Duble  :", i)
			return false
		}
	}
	return true
}
