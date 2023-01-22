package main

import (
	"bytes"
	"image"
	"image/jpeg"
	"log"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/storage"
)

type bgImageLoad struct {
	uri fyne.URI
	img *canvas.Image
}

var uriarray = []fyne.URI{
	storage.NewURI("https://cdn.fishki.net/upload/post/2023/01/20/4342685/8cf6c3d577fef2bfa31987d1ff1c7798.jpg"),
	storage.NewURI("https://cdn.fishki.net/upload/post/2023/01/20/4342685/75fd75e212a88ae8138cad4834ae6d06.jpg"),
	storage.NewURI("https://cdn.fishki.net/upload/post/2023/01/20/4342685/5efeb0a414b0a0db0725302c5ecfd7b9.jpg"),
}

var loads = make(chan bgImageLoad, 1024)

// var w2 fyne.Window

func loadImage(u fyne.URI) fyne.CanvasObject {
	img := canvas.NewImageFromResource(nil)
	img.FillMode = canvas.ImageFillContain
	img.SetMinSize(fyne.Size{320, 320})

	loads <- bgImageLoad{u, img}
	return img
}

func doLoadImage(u fyne.URI, img *canvas.Image) {
	read, err := storage.OpenFileFromURI(u)
	if err != nil {
		log.Println("Error opening image", err)
		return
	}
	defer read.Close()
	raw, _, err := image.Decode(read)
	if err != nil {
		log.Println("Error decoding image", err)
		return
	}
	img.Image = raw
	img.Refresh()
	wr, err := storage.Writer(storage.NewFileURI("c:/img/" + u.String()[strings.LastIndex(u.String(), "/"):]))
	defer wr.Close()
	buf := new(bytes.Buffer)
	jpeg.Encode(buf, img.Image, nil)
	wr.Write(buf.Bytes())
	// fmt.Println(" Loaded and refreshed ....")
	w.Content().Refresh()
}

func doLoadImages() {
	// fmt.Println("Start routine ... .")
	for load := range loads {
		// fmt.Println("Cycling inside ...")
		// fmt.Println("load.uri -> ", load.uri, "load.img -> ", load.img)

		doLoadImage(load.uri, load.img)
	}
}

// func scaleImage(img image.Image) image.Image {
// 	return resize.Thumbnail(320, 240, img,
// 		resize.Lanczos3)
// }

// func main() {

// 	appm := app.New()
// 	w2 = appm.NewWindow("App")
// 	contentall := container.NewGridWrap(fyne.Size{380, 380}, loadImage(uriarray[0]), loadImage(uriarray[1]), loadImage(uriarray[2]))
// 	w2.SetContent(contentall)
// 	w2.Resize(fyne.Size{980, 720})
// 	w2.RequestFocus()
// 	w2.Content().Refresh()
// 	go doLoadImages()
// 	w2.ShowAndRun()
// }
