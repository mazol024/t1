package main

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	// "fyne.io/fyne/v2/internal/widget"
	// "fyne.io/fyne/v2/internal/widget"
)

// type wdlbtype widget.Button

// // type wdlbtype canvas.Circle

// func (t wdlbtype) FocusGained() { fmt.Println("Ganed focus") }
// func (t wdlbtype) FocusLost() {
// 	log.Println("Focuse lost")
// 	os.Exit(0)
// }
// func (t wdlbtype) TypedRune(r rune)        { fmt.Println(strconv.QuoteRune(r)) }
// func (t wdlbtype) TypedKey(*fyne.KeyEvent) {}

type t struct {
	widget.Button
}

func (t *t) FocusGained() {
	log.Println("FOcus", t.Button.Text)

}
func (t *t) FocusLost() {
	log.Println("FOcus Lost", t.Button.Text)
	// (*widget.Button)(&tb1.Button).Text = "New Text3"
	contt.Refresh()
}

func (t *t) TypedRune(r rune) {
	fmt.Println(strconv.QuoteRune(r))
}
func (t *t) TypedKey(k *fyne.KeyEvent) {
	fmt.Println(k.Name)
}

type p struct {
	*widget.ButtonStyle
}

var tb1 *t
var tb2 *t
var tb3 *t
var contt *fyne.Container

func main() {
	myApp := app.New()
	// var imgarr []fyne.CanvasObject
	myWindow := myApp.NewWindow("WD Focused")
	tb0 := widget.NewButton("Non focused ", func() {
		log.Println("tapped")
	})

	tb0.Resize(fyne.Size{20, 20})
	tb1 = &t{*widget.NewButton("Button N1", func() {
		log.Println("tapped")
	}),
	}
	tb2 = &t{*widget.NewButton("Button N2 ", func() {
		log.Println("tapped")
	}),
	}
	tb3 = &t{*widget.NewButton("Button N3 ", func() {
		log.Println("tapped")
	}),
	}
	tb1.Importance = widget.MediumImportance
	tb1.Resize(fyne.Size{20, 20})
	tb1.Refresh()
	// conttt := fyne.NewContainer(tb1)
	// circle.StrokeColor = color.Gray{0x99}
	// circle.Resize(fyne.Size{200, 200})
	// btn1 = (*wdlbtype)(t)
	// btn1.Resize(fyne.Size{120, 120})
	// *&btn1.Text = " text new "
	contt = container.New(layout.NewVBoxLayout(), tb0, tb1, tb2, tb3)
	// cont := container.New(layout.NewVBoxLayout(), t, (*widget.Button)(t2), t2)
	// cont := container.New(layout.NewMaxLayout(), (btn1))
	// fyne.Canvas.Focus(myWindow.Canvas(), btn1)
	go func() {
		time.Sleep(time.Second * 3)
		tb0.SetText("new Text")
		tb0.Text = "new Botton 0 text"
		// tb1.SetText("new Text")
		// (*widget.Button)(&tb1.Button).Text = "New Text2"
		tb1.Button.Refresh()
		(*widget.Button)(&tb1.Button).Refresh()
	}()
	myWindow.SetContent(contt)
	myWindow.Resize(fyne.Size{600, 400})
	myWindow.ShowAndRun()
}

//	func scaleImage(img image.Image) image.Image {
//		return resize.Thumbnail(320, 240, img,
//			resize.Lanczos3)
//	}
