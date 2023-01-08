package main

import (
	"image/color"
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Layout")

	str := binding.NewString()
	str.Set("Initial value")

	label1 := widget.NewLabelWithData(str)
	input := widget.NewEntry()
	input.SetPlaceHolder("Enter text...")
	img := canvas.NewImageFromFile("1.png")
	img1 := canvas.NewImageFromFile("1.png")
	// img.FillMode = canvas.ImageFillOriginal
	// img1.FillMode = canvas.ImageFillOriginal
	button := widget.NewButton("click me", func() {
		log.Println("tapped")
		p := input.Text
		str.Set(p)
	})
	button1 := widget.NewButton("Clear Input", func() {
		input.SetText("Initial value")
		str.Set("Initial value")
	})
	text0 := canvas.NewText("Continue", color.RGBA{251, 123, 228, 255})
	text0.TextSize = 32
	text0.TextStyle = fyne.TextStyle{Bold: true, Italic: true}
	text1 := canvas.NewText("Second Device", color.Black)
	text2 := canvas.NewText("First device", color.Black)

	cont1 := container.New(layout.NewVBoxLayout(), input, label1)
	cont2 := container.New(layout.NewHBoxLayout(), button, button1)
	content := container.New(layout.NewVBoxLayout(), cont1, cont2, text2, img1, text0, text1, img)
	text0.SetMinSize(fyne.Size{30, 30})
	text1.SetMinSize(fyne.Size{10, 10})
	text2.SetMinSize(fyne.Size{10, 10})
	img.SetMinSize(fyne.Size{260, 260})
	img1.SetMinSize(fyne.Size{260, 260})
	// img.Resize(fyne.Size{160, 160})
	// img1.Resize(fyne.Size{160, 160})
	myWindow.SetContent(content)

	myWindow.Resize(fyne.Size{159, 200})
	myWindow.ShowAndRun()
}
