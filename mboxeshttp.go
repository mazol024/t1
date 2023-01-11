package main

import (
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"
)

type CustomImagem struct {
	widget.Card
	link1 string
}

func (img *CustomImagem) Tapped(ev *fyne.PointEvent) {
	// fmt.Println("Tapped x= ", ev.AbsolutePosition.X, "  y= ", ev.AbsolutePosition.Y)
	// fmt.Println("You choose: ", img.Subtitle)
	showPic(img.link1)
	// showPic(img.Subtitle)
	// w.Canvas().Content().Refresh()
	// contentall.Refresh()
	// contentall.Show()
}
func (img *CustomImagem) MouseIn(*desktop.MouseEvent) {
	// fmt.Println("Entered", img.Title, img.Subtitle)
}

func (img *CustomImagem) MouseOut() {
	// fmt.Println("Exited")
}

func (img *CustomImagem) MouseMoved(*desktop.MouseEvent) {

}

var appm fyne.App
var w fyne.Window
var imgp []fyne.CanvasObject
var img1 *canvas.Image
var contentm *CustomImagem
var contentall *fyne.Container

// var imgp []CustomImagem

func drawinghttp(images []string, site string) {
	appm = app.New()
	w = appm.NewWindow("test-image")

	for n, i := range images {
		img1 = canvas.NewImageFromURI(storage.NewURI(i))
		// img1.FillMode = canvas.ImageFillOriginal
		img1.SetMinSize(fyne.Size{64, 64})
		// ttext := canvas.NewText("hello", color.Black)
		box := widget.NewCard("", "#"+strconv.Itoa(n), img1)
		contentm = &CustomImagem{*box, i}
		// contentm.Resize(fyne.Size{320, 320})
		imgp = append(imgp, contentm)
		if n == 30 {
			break
		}
	}

	// box := widget.NewHBox(image)
	// text1 := widget.NewLabel("Hello ")
	// contentm := container.New(layout.NewHBoxLayout(), text1)
	// contentall = container.New(layout.NewGridLayout(4), imgp[:]...)
	contentall = container.NewGridWrap(fyne.Size{320, 320}, imgp[:]...)
	scroll1 := container.NewScroll(contentall)

	w.SetContent(scroll1)
	w.Resize(fyne.Size{980, 720})
	w.RequestFocus()
	w.ShowAndRun()
}
func showPic(i string) {
	// ii := canvas.NewImageFromFile("1.png")
	// ii := i.Image
	// p := container.New(layout.NewMaxLayout(), i.Image)

	// p.Resize(fyne.Size{220, 220})
	// p.FillMode = canvas.ImageFillOriginal
	// i.Image.SetMinSize(fyne.Size{600, 600})
	// d := dialog.NewCustom("Picture", "Close", &(*i.Image), w)
	img2 := canvas.NewImageFromURI(storage.NewURI(i))
	d := dialog.NewCustom("Picture", "Close", img2, w)
	// i.Image.FillMode = canvas.ImageFillContain
	img2.FillMode = canvas.ImageFillOriginal

	d.Resize(fyne.Size{1200, 980})
	// d := dialog.NewCustom("Picture", "Close", p, w)
	// d.SetOnClosed(w.Content().Refresh)
	// d.SetOnClosed(w.Canvas().Content().Refresh)

	d.Show()
	// w.SetContent(contentall)
	// w.Resize(fyne.Size{980, 720})
	w.RequestFocus()
	// contentall.Refresh()
	// contentall.Show()
}
