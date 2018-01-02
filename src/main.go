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
var blue = color.RGBA{0, 0, 255, 255}

func main() {
	img := image.NewRGBA(image.Rect(0, 0, 255, 255))

	// Set Color point. Temporary. Only for test.
	// img.Set(52, 41, red)

	img = library.FlipByVertically(image.RGBA{Pix: img.Pix, Stride: img.Stride, Rect: img.Rect})

	newImage := createLine(13, 20, 80, 40, image.RGBA{Pix: img.Pix, Stride: img.Stride, Rect: img.Rect}, red)
	newImage = createLine(20, 13, 40, 80, image.RGBA{Pix: newImage.Pix, Stride: newImage.Stride, Rect: newImage.Rect}, white)
	newImage = createLine(80, 40, 13, 20, image.RGBA{Pix: newImage.Pix, Stride: newImage.Stride, Rect: newImage.Rect}, blue)



	//TODO: Make as separate func in library.
	f, _ := os.OpenFile("out.png", os.O_WRONLY|os.O_CREATE, 0600)
	defer f.Close()
	png.Encode(f, newImage.SubImage(newImage.Rect))

}

// Return value only for dev test.
func createLine(x0, y0, x1, y1 int, img image.RGBA, color color.RGBA) image.RGBA {
	for i := x0; i <= x1; i++ {
		x := i - x0 / x1 - x0
		y := y0*(1 - x) + y1*x
		img.Set(int(x), int(y), color)
	}
	return img
}
