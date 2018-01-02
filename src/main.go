package main

import (
	// "AuroraRender/library"
	"image"
	"image/color"
	"os"
	"image/png"
	"math"
)

//TODO: Make as const in future release.
var white = color.RGBA{255, 255, 255, 255}
var red = color.RGBA{255, 0, 0, 255}
var blue = color.RGBA{0, 0, 255, 255}

func main() {
	img := image.NewRGBA(image.Rect(0, 0, 255, 255))

	// Set Color point. Temporary. Only for test.
	// img.Set(52, 41, red)

	// Incorrect flip. Need fix.
	// img = library.FlipByVertically(image.RGBA{Pix: img.Pix, Stride: img.Stride, Rect: img.Rect})

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
	steep := false

	if math.Abs(float64(x0 - x1)) < math.Abs(float64(y0 - y1)) {
		x0, y0 = y0, x0
		x1, y1 = y1, x1
		steep = true
	}

	if x0 > x1 {
		x0, x1 = x1, x0
		y0, y1 = y1, y0
	}

	dx := x1 - x0
	dy := y1 - y0

	derror := math.Abs(float64(dy)) * 2
	error := 0.0

	y := y0

	for x := x0; x <= x1; x++ {
		if steep {
			img.Set(int(y), int(x), color)
		} else {
			img.Set(int(x), int(y), color)
		}

		error += derror

		if error > float64(dx) {
			if y1 > y0 { y += 1 } else { y += -1 }
			error -= float64(dx * 2)
		}

	}
	return img
}
