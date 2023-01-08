package main

import (
	"image/color"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
)

func main() {
	myApp := app.New()
	w := myApp.NewWindow("Raster")

	// raster := canvas.NewRasterWithPixels(
	// 	func(_, _, w, h int) color.Color {
	// 		return color.RGBA{uint8(rand.Intn(255)),
	// 			uint8(rand.Intn(255)),
	// 			uint8(rand.Intn(255)), 0xff}
	// 	})
	image := canvas.NewImageFromFile("1.png")
	image1 := canvas.NewImageFromFile("1.png")
	image.FillMode = canvas.ImageFillOriginal
	image1.FillMode = canvas.ImageFillOriginal
	// image.SetMinSize(fyne.Size{200, 200})
	image.SetMinSize(fyne.Size{100, 200})
	image1.SetMinSize(fyne.Size{100, 100})
	// Next line
	// text := canvas.NewText("Text Object", color.Black)
	// raster.Resize(fyne.NewSize(120, 100))

	circle := canvas.NewCircle(color.RGBA{135, 35, 222, 255})
	circle.StrokeColor = color.Gray{0x99}
	circle.StrokeWidth = 5
	circle.Resize(fyne.NewSize(80, 80))

	// content0 := container.New(layout.NewHBoxLayout(), circle)
	// content0.Resize(fyne.Size{200, 200})
	text0 := canvas.NewText("Text", color.RGBA{128, 128, 226, 255})
	text1 := canvas.NewText("Text1", color.RGBA{128, 128, 226, 255})
	text2 := canvas.NewText("Text2", color.RGBA{128, 92, 226, 255})

	content1 := container.New(layout.NewHBoxLayout(), text0, text1, text2, image, image1)

	text3 := canvas.NewText("Text3", color.RGBA{98, 222, 226, 255})
	text4 := canvas.NewText("Text4", color.RGBA{222, 222, 226, 255})
	text5 := canvas.NewText("Text5", color.RGBA{253, 255, 56, 255})
	content2 := container.New(layout.NewHBoxLayout(), text4, text3, text5, image, image1)
	content := container.New(layout.NewVBoxLayout(), content1, content2)
	content11 := container.NewBorder(nil, content, nil, nil, image, circle)
	w.SetContent(content11)
	w.Resize(fyne.NewSize(820, 640))
	go func() {
		time.Sleep(time.Second * 4)
		circle.Hide()
		time.Sleep(time.Second * 4)
		image.SetMinSize(fyne.Size{300, 300})
		image1.Resize(fyne.Size{600, 300})
		image1.Refresh()
		image.Refresh()
	}()
	w.ShowAndRun()
}
