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

func (img *CustomImagem) Tapped(ev *fyne.PointEvent) {
	fmt.Println("Tapped x= ", ev.AbsolutePosition.X, "  y= ", ev.AbsolutePosition.Y)
	fmt.Println("You choose: ", img.Subtitle)

}
func (img *CustomImagem) MouseIn(*desktop.MouseEvent) {
	fmt.Println("Entered", img.Title, img.Subtitle)
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
		img.SetMinSize(fyne.Size{72, 72})
		img.FillMode = canvas.ImageFillOriginal
		box := widget.NewCard("Image #"+strconv.Itoa(n), i, img)
		contentm := &CustomImagem{*box}
		imgp = append(imgp, contentm)
		if n == 25 {
			break
		}
	}

	// box := widget.NewHBox(image)
	// text1 := widget.NewLabel("Hello ")
	// contentm := container.New(layout.NewHBoxLayout(), text1)
	contentall := container.New(layout.NewGridLayout(4), imgp[:]...)
	w.SetContent(contentall)
	w.Resize(fyne.Size{980, 720})
	w.RequestFocus()
	w.ShowAndRun()
}
