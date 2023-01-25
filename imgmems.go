package main

import (
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/layout"
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
	showPic(img)
}
func (img *CustomImagem) MouseIn(*desktop.MouseEvent) {
	// fmt.Println("Entered", img.Title, img.Subtitle)
}

func (img *CustomImagem) MouseOut() {
	// fmt.Println("Exited")
}

func (img *CustomImagem) MouseMoved(*desktop.MouseEvent) {

}

var images []string
var appm fyne.App
var w fyne.Window

// var imgp []CustomImagem

var imgp []fyne.CanvasObject
var img1 *canvas.Image
var contentm *CustomImagem
var scroll1 *container.Scroll
var contentall *fyne.Container
var contentlist *fyne.Container
var contenLast *container.Split

var indexS string

func main() {
	appm = app.New()
	w = appm.NewWindow("App")

	urls := getThemes()
	keys := []string{}
	for k, _ := range urls {
		keys = append(keys, k)
	}
	selectEntry1 := widget.NewSelect(keys, func(s string) {
		indexS = urls[s]

		images := runHttp3(indexS)
		imgp = []fyne.CanvasObject{}
		for n, i := range images {
			box := widget.NewCard("", "#"+strconv.Itoa(n), loadImage(storage.NewURI(i)))
			box.Resize(fyne.Size{320, 320})
			contentm := &CustomImagem{*box, i}
			// contentm.Resize(fyne.Size{320, 320})
			imgp = append(imgp, contentm)
		}
		contentall = container.NewGridWrap(fyne.Size{160, 160}, imgp[:]...)
		scroll1 = container.NewScroll(contentall)
		scroll1.SetMinSize(fyne.Size{980, 640})
		contenLast = container.NewVSplit(contentlist, scroll1)
		w.SetContent(contenLast)
		w.Content().Refresh()
	})

	contentlist = container.New(layout.NewVBoxLayout(), selectEntry1)
	contentall = container.NewGridWrap(fyne.Size{160, 160})
	scroll1 = container.NewScroll(contentall)
	scroll1.SetMinSize(fyne.Size{980, 640})
	contenLast = container.NewVSplit(contentlist, scroll1)
	w.SetContent(contenLast)
	w.Resize(fyne.Size{980, 720})
	w.Content().Refresh()
	w.RequestFocus()
	// doLoadImages()
	go doLoadImages()
	w.ShowAndRun()

}
func showPic(img24 *CustomImagem) {
	// func showPic(img24 *CustomImagem) {
	vv := img24
	// vv := &img24.Card
	pos := vv.Position()
	ss := vv.Size()
	// bigimg := container.NewCenter(img24)
	bigimg := container.NewMax(img24)
	bigimg.Resize(fyne.Size{980, 720})
	// img24.Resize(fyne.Size{1980, 1020})
	d := dialog.NewCustom("", "Close", bigimg, w)
	// d := dialog.NewCustom("", "Close", img24, w)
	d.Resize(fyne.Size{980, 720})
	d.Show()
	// d.Refresh()
	// w.Content().Refresh()
	d.SetOnClosed(func() {
		vv.Resize(ss)
		vv.Move(pos)
		contentall = container.NewGridWrap(fyne.Size{160, 160}, imgp[:]...)
		scroll1 = container.NewScroll(contentall)
		scroll1.SetMinSize(fyne.Size{980, 640})
		contenLast = container.NewVSplit(contentlist, scroll1)
		w.SetContent(contenLast)
		// w.Content().Refresh()
	})
}
