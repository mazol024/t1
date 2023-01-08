package main

import (
	"fmt"
	"log"
	"strconv"

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

type pp widget.Label

// type pp fyne.CanvasObject
// type pp canvas.Image

func (t *pp) FocusGained()            { fmt.Println("Ganed focus") }
func (t *pp) FocusLost()              { log.Println("Focuse lost") }
func (t *pp) TypedRune(r rune)        { fmt.Println(strconv.QuoteRune(r)) }
func (t *pp) TypedKey(*fyne.KeyEvent) {}

var lab, btn *pp

// var imgpp []pp

var imgpp []fyne.CanvasObject

var img *canvas.Image

func printimg(images []string, site string) {
	myApp := app.New()
	// var imgarr []fyne.CanvasObject
	myWindow := myApp.NewWindow("Pictures from site" + site)

	for n, i := range images {
		img := canvas.NewImageFromURI(storage.NewURI(i))
		// ppr = canvas.NewImageFromURI(storage.NewURI(i))
		img.SetMinSize(fyne.Size{170, 120})
		// pp.FocusGained()
		// imgarr = append(imgarr, sim{img})
		imgpp = append(imgpp, img)
		if n == 10 {
			break
		}
	}

	var content *fyne.Container
	// content = container.New(layout.NewGridLayout(5), *imgpp[:]...)
	lab = (*pp)(widget.NewLabel("test label here "))
	btn = (*pp)(widget.NewLabel("test label here "))

	// lab = lab1
	content = container.New(layout.NewGridLayout(5), imgpp[:]...)
	cont := container.New(layout.NewGridLayout(2), btn, lab, content)

	// content0 := container.NewVScroll(lab, content)
	myWindow.SetContent(cont)
	myWindow.Canvas().Focus(lab)
	myWindow.Resize(fyne.Size{1200, 740})
	myWindow.ShowAndRun()
}

//	func scaleImage(img image.Image) image.Image {
//		return resize.Thumbnail(320, 240, img,
//			resize.Lanczos3)
//	}
