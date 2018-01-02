package main

import (
	"AuroraRender/library"
	"image"
	"image/color"
	"os"
	"image/png"
)

//TODO: Make as const in future release.
var white = color.RGBA{255, 255, 255, 255}
var red = color.RGBA{255, 0, 0, 255}

func main() {
	img := image.NewRGBA(image.Rect(0, 0, 100, 100))

	img.Set(52, 41, red)

	img = library.FlipByVertically(image.RGBA{Pix: img.Pix, Stride: img.Stride, Rect: img.Rect})

	//TODO: Make as separate func in library.
	f, _ := os.OpenFile("out.png", os.O_WRONLY|os.O_CREATE, 0600)
	defer f.Close()
	png.Encode(f, img)

}
