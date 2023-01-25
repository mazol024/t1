package main

import (
	"image"
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/storage"
)

type bgImageLoad struct {
	uri fyne.URI
	img *canvas.Image
}

var loads = make(chan bgImageLoad, 2048)

func loadImage(u fyne.URI) fyne.CanvasObject {
	img := canvas.NewImageFromResource(nil)
	img.FillMode = canvas.ImageFillContain
	img.SetMinSize(fyne.Size{480, 480})

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
	// fmt.Println(" Loaded and refreshed ....")
}

func doLoadImages() {
	// fmt.Println("Start routine ... .")
	for load := range loads {
		// fmt.Println("Cycling inside ...")
		// fmt.Println("load.uri -> ", load.uri, "load.img -> ", load.img)
		doLoadImage(load.uri, load.img)

	}
	w.Content().Refresh()
}
