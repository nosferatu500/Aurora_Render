package main

import (
	"image"
	"image/color"
	"os"
	"image/png"
	"math"
	"AuroraRender/library"
	"fmt"
	"strconv"
)

const (
	width = 800
	height = 800
	depth = 255
)

//TODO: Make as const in future release.
var white = color.RGBA{255, 255, 255, 255}
var red = color.RGBA{255, 0, 0, 255}
var blue = color.RGBA{0, 0, 255, 255}

func main() {
	img := image.NewRGBA(image.Rect(0, 0, width, height))

	var newImage image.RGBA

	var zBuffer [width*height]int

	for i:=0; i<width*height; i++ {
		zBuffer[i] = -2147483648
	}

	model := library.CreateModel("./obj/african_head.obj")

	count := library.GetFaceCount(model)

	// Fix bug with out of range for faces collection.
	for i := 0; i < count - 20; i++ {
		face := model.Faces[i]

		var screenCoords [3]library.Vector2D
		var worldCoords [3]library.Vector3D

		for j := 0; j < 3; j++ {
			v := model.Verts[face[j] - 1]
			screenCoords[j] = library.Vector2D{(v.X + 1) * width / 2, (v.Y + 1) * height / 2}
			worldCoords[j] = v
		}

		edge1 := worldCoords[2].Subtract(worldCoords[0])
		edge2 := worldCoords[1].Subtract(worldCoords[0])

		n := edge1.Multiply(edge2)

		n = n.Normalize()

		intensity := n.Multiply(library.Vector3D{0,0,-1})

		// newImage = createLine(screenCoords[1], screenCoords[2], image.RGBA{Pix: img.Pix, Stride: img.Stride, Rect: img.Rect}, white)

		//r := rand.New(rand.NewSource(int64(i)))
		//randomColor := color.RGBA{uint8(r.Int()), uint8(r.Int()), uint8(r.Int()), 255}
		//newImage = createTriangle(screenCoords[0], screenCoords[1], screenCoords[2],image.RGBA{Pix: img.Pix, Stride: img.Stride, Rect: img.Rect}, randomColor)


		if intensity.Z > 0 {
			newNumberString := strconv.FormatFloat(intensity.Z * 255, 'f', 0, 64)
			newNumber, _ := strconv.ParseInt(newNumberString, 10, 8)
			fmt.Println(newNumber)

			intensityColor := color.RGBA{uint8(newNumber), uint8(newNumber), uint8(newNumber), 255}
			newImage = createTriangle(screenCoords[0], screenCoords[1], screenCoords[2],image.RGBA{Pix: img.Pix, Stride: img.Stride, Rect: img.Rect}, intensityColor)
		}

	}

	img = library.FlipByVertically(image.RGBA{Pix: newImage.Pix, Stride: newImage.Stride, Rect: newImage.Rect})

	//TODO: Make as separate func in library.
	f, _ := os.OpenFile("out.png", os.O_WRONLY|os.O_CREATE, 0600)
	defer f.Close()
	png.Encode(f, img)

}

func createLine(p0, p1 library.Vector2D, img image.RGBA, color color.RGBA) image.RGBA {
	steep := false

	if math.Abs(float64(p0.X - p1.X)) < math.Abs(float64(p0.Y - p1.Y)) {
		p0.X, p0.Y = p0.Y, p0.X
		p1.X, p1.Y = p1.Y, p1.X
		steep = true
	}

	if p0.X > p1.X {
		p0, p1 = p1, p0
	}

	for x := p0.X; x <= p1.X; x++ {
		t := (x-p0.X)/(p1.X-p0.X);
		y := p0.Y*(1.-t) + p1.Y*t;
		if steep {
			img.Set(int(y), int(x), color);
		} else {
			img.Set(int(x), int(y), color);
		}
	}
	return img
}

func createTriangle(t0, t1, t2 library.Vector2D, img image.RGBA, color color.RGBA) image.RGBA {
	if t0.Y > t1.Y { t0, t1 = t1, t0 }

	if t0.Y > t2.Y { t0, t2 = t2, t0 }

	if t1.Y > t2.Y { t1, t2 = t2, t1 }

	totalHeight := t2.Y - t0.Y

	for y := t0.Y; y <= t1.Y; y++ {
		segmentHeight := t1.Y-t0.Y+1
		alpha := (y-t0.Y) / totalHeight
		beta := (y-t0.Y) / segmentHeight

		firstStepA := t2.Subtract(t0)
		secondStepA := firstStepA.MultiplyScalar(alpha)
		A := t0.Add(secondStepA)

		firstStepB := t1.Subtract(t0)
		secondStepB := firstStepB.MultiplyScalar(beta)
		B := t0.Add(secondStepB)

		if A.X > B.X { A, B = B, A }

		for j := A.X; j <= B.X; j++ {
			img.Set(int(j), int(y), color)
		}
	}

	for y := t1.Y; y <= t2.Y; y++ {
		segmentHeight := t2.Y-t1.Y+1
		alpha := (y-t0.Y) / totalHeight
		beta := (y-t1.Y) / segmentHeight

		firstStepA := t2.Subtract(t0)
		secondStepA := firstStepA.MultiplyScalar(alpha)
		A := t0.Add(secondStepA)

		firstStepB := t2.Subtract(t1)
		secondStepB := firstStepB.MultiplyScalar(beta)
		B := t1.Add(secondStepB)

		if A.X > B.X { A, B = B, A }

		for j := A.X; j <= B.X; j++ {
			img.Set(int(j), int(y), color)
		}
	}
	return img
}
