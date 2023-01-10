package main

import (
	"fmt"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"
)

type CustomImagem struct {
	widget.Card
}

func (img *CustomImagem) MouseIn(*desktop.MouseEvent) {
	fmt.Println("Entered", img.Title, img.Subtitle)
	img.Title = "Choosen "
	w.Canvas().Refresh(img)
}

func (img *CustomImagem) MouseOut() {
	fmt.Println("Exited")
}

func (img *CustomImagem) MouseMoved(*desktop.MouseEvent) {

}

var appm fyne.App
var w fyne.Window
var imgp []fyne.CanvasObject

// var imgp []CustomImagem

func drawinghttp(images []string, site string) {
	appm = app.New()
	w = appm.NewWindow("test-image")

	for n, i := range images {
		img := canvas.NewImageFromURI(storage.NewURI(i))
		// image := &canvas.Image{
		// 	File:     "1.png",
		// 	FillMode: canvas.ImageFillOriginal,
		// }
		// ppr = canvas.NewImageFromURI(storage.NewURI(i))
		// image.SetMinSize(fyne.NewSize(280, 280))
		img.SetMinSize(fyne.Size{64, 64})
		box := widget.NewCard("Image #"+strconv.Itoa(n), i, img)
		contentm := &CustomImagem{*box}
		// pp.FocusGained()
		contentm.Resize(fyne.Size{64, 64})
		// imgarr = append(imgarr, sim{img})
		imgp = append(imgp, contentm)
		if n == 25 {
			break
		}
	}

	// box := widget.NewHBox(image)
	// text1 := widget.NewLabel("Hello ")
	// contentm := container.New(layout.NewHBoxLayout(), text1)
	contentall := container.New(layout.NewGridLayout(5), imgp[:]...)
	contentall.Resize(fyne.Size{1200, 900})
	w.SetContent(contentall)
	w.Resize(fyne.Size{980, 720})
	w.RequestFocus()
	w.ShowAndRun()
}
