package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/widget"
)

type CustomImage struct {
	widget.Card
}

func (img *CustomImage) MouseIn(*desktop.MouseEvent) {
	fmt.Println("Entered", img.Title, img.Subtitle)
}

func (img *CustomImage) MouseOut() {
	fmt.Println("Exited")
}

func (img *CustomImage) MouseMoved(*desktop.MouseEvent) {

}

func main() {
	app := app.New()

	w := app.NewWindow("test-image")

	image := &canvas.Image{
		File:     "1.png",
		FillMode: canvas.ImageFillOriginal,
	}
	image.SetMinSize(fyne.NewSize(280, 280))

	box := widget.NewCard("File name :  ", "Sub", image)
	// box := widget.NewHBox(image)
	contentm := &CustomImage{*box}
	// text1 := widget.NewLabel("Hello ")
	// contentm := container.New(layout.NewHBoxLayout(), text1)

	w.SetContent(contentm)

	w.RequestFocus()
	w.ShowAndRun()
}
