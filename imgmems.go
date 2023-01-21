package main

import (
	"fmt"

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

var images []string
var appm fyne.App
var w fyne.Window
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
		fmt.Println("Choosen", indexS)

		images := runHttp3(indexS)
		w.Content().Refresh()

		imgp := []fyne.CanvasObject{}
		for n, i := range images {
			imgp = append(imgp, loadImage(storage.NewURI(i)))
			if n == 20 {
				break
			}
		}
		contentall = nil
		contentall = container.NewGridWrap(fyne.Size{320, 320}, imgp[:]...)
		scroll1 = container.NewScroll(contentall)
		scroll1.SetMinSize(fyne.Size{980, 640})
		contenLast = nil
		contenLast = container.NewVSplit(contentlist, scroll1)
		contentlist.Show()
		w.SetContent(contenLast)
		w.Content().Refresh()
	})

	contentlist = container.New(layout.NewVBoxLayout(), selectEntry1)
	contentall = container.NewGridWrap(fyne.Size{320, 320})
	scroll1 = container.NewScroll(contentall)
	scroll1.SetMinSize(fyne.Size{980, 640})
	contenLast = container.NewVSplit(contentlist, scroll1)
	contentlist.Show()
	w.SetContent(contenLast)

	w.Resize(fyne.Size{980, 720})
	w.RequestFocus()
	w.Content().Refresh()
	doLoadImages()
	w.ShowAndRun()

}
func showPic22(i string) {
	img2 := canvas.NewImageFromURI(storage.NewURI(i))
	d := dialog.NewCustom(i, "Close", img2, w)
	img2.FillMode = canvas.ImageFillOriginal
	d.Resize(fyne.Size{1200, 980})
	d.Show()
	w.RequestFocus()
}
