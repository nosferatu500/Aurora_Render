package main

import (
	"image"
	"image/color"
	"os"
	"image/png"
	"math"
	"AuroraRender/library"
	"fmt"
)

const (
	width = 800
	height = 800
)

//TODO: Make as const in future release.
var white = color.RGBA{255, 255, 255, 255}
var red = color.RGBA{255, 0, 0, 255}
var blue = color.RGBA{0, 0, 255, 255}

func main() {
	img := image.NewRGBA(image.Rect(0, 0, width, height))

	var newImage image.RGBA

	model := library.CreateModel("./obj/african_head.obj")

	count := library.GetFaceCount(model)

	// Fix bug with out of range for faces collection.
	for i := 0; i < count - 300; i++ {
		face := model.Faces[i]

		fmt.Println(i, face, count)

		for j := 0; j < 3; j++ {
			v0 := model.Verts[face[j]]
			v1 := model.Verts[face[(j + 1) % 3]]

			x0 := (v0.X + 1.0) * width / 2
			y0 := (v0.Y + 1.0) * height / 2

			x1 := (v1.X + 1.0) * width / 2
			y1 := (v1.Y + 1.0) * height / 2

			newImage = createLine(int(x0), int(y0), int(x1), int(y1), image.RGBA{Pix: img.Pix, Stride: img.Stride, Rect: img.Rect}, white)
		}
	}

	img = library.FlipByVertically(image.RGBA{Pix: newImage.Pix, Stride: newImage.Stride, Rect: newImage.Rect})

	//TODO: Make as separate func in library.
	f, _ := os.OpenFile("out.png", os.O_WRONLY|os.O_CREATE, 0600)
	defer f.Close()
	png.Encode(f, img)

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
