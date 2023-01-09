package main

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"

	// "fyne.io/fyne/v2/internal/widget"

	"fyne.io/fyne/v2/widget"

	// "fyne.io/fyne/v2/internal/widget"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/storage"
)

// type imgimg *canvas.Image

type pp struct {
	*widget.Label
}

// type pp fyne.CanvasObject
// type pp canvas.Image

func (t pp) FocusGained() {
	fmt.Println("Ganed focus")
	// t.TextStyle.Bold = true
	// t.TextStyle.Italic = false
	(t.Label).SetText(" New text ")

}
func (t pp) FocusLost() {
	log.Println("Focuse lost")
	t.TextStyle.Italic = true
	t.TextStyle.Bold = false

	p := cont.Objects
	p[1].Show()
	p[0].Show()
	p[2].Show()
	p[3].Show()
	(t).SetText(" Old  text ")
	(&t).SetText(" tpooo Old  text ")
	content.Refresh()
	// myWindow.Content().Refresh()
}
func (t *pp) TypedRune(r rune)        { fmt.Println(strconv.QuoteRune(r)) }
func (t *pp) TypedKey(*fyne.KeyEvent) {}

var content, cont *fyne.Container

var lab, btn *pp

// var imgpp []pp

var imgpp []fyne.CanvasObject

var img *canvas.Image
var myWindow fyne.Window

func printimg(images []string, site string) {
	myApp := app.New()
	// var imgarr []fyne.CanvasObject
	myWindow = myApp.NewWindow("Pictures from site" + site)

	for n, i := range images {
		img := canvas.NewImageFromURI(storage.NewURI(i))
		// ppr = canvas.NewImageFromURI(storage.NewURI(i))
		img.SetMinSize(fyne.Size{64, 64})
		// pp.FocusGained()
		// imgarr = append(imgarr, sim{img})
		imgpp = append(imgpp, img)
		if n == 10 {
			break
		}
	}

	// content = container.New(layout.NewGridLayout(5), *imgpp[:]...)
	lab = &pp{widget.NewLabel("First label here ")}

	btn = &pp{widget.NewLabel("Second label here ")}
	tri := widget.NewLabel("Third label here ")

	// lab = lab1
	content = container.New(layout.NewGridLayout(5), imgpp[:]...)
	cont = container.New(layout.NewVBoxLayout(), btn, lab, content, tri)
	go func() {
		time.Sleep(time.Second * 8)
		tri.SetText("New Taexte ")
		lab.TextStyle.Bold = true
		lab.TextStyle.Italic = true
		// *lab.Text = "   "
		(lab.Label).SetText("Text Changed ")
		(btn.Label).SetText("New Tesxt here ")
		cont.Refresh()
		// myWindow.Content().Refresh()
	}()
	// content0 := container.NewVScroll(lab, content)
	myWindow.SetContent(cont)
	// myWindow.Canvas().Focus(btn)
	myWindow.Resize(fyne.Size{1200, 740})
	myWindow.RequestFocus()
	myWindow.ShowAndRun()

}

//	func scaleImage(img image.Image) image.Image {
//		return resize.Thumbnail(320, 240, img,
//			resize.Lanczos3)
//	}
