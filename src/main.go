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

		fmt.Println(i, face, count)

		for j := 0; j < 3; j++ {
			v0 := model.Verts[face[j] - 1]
			v1 := model.Verts[face[(j+1)%3] - 1]

			x0 := int((v0.X + 1) * float64(width) / 2)
			y0 := int((v0.Y + 1) * float64(height) / 2)

			x1 := int((v1.X + 1) * float64(width) / 2)
			y1 := int((v1.Y + 1) * float64(height) / 2)

			newImage = createLine(x0, y0, x1, y1, image.RGBA{Pix: img.Pix, Stride: img.Stride, Rect: img.Rect}, white)
		}
	}

	img = library.FlipByVertically(image.RGBA{Pix: newImage.Pix, Stride: newImage.Stride, Rect: newImage.Rect})

	//TODO: Make as separate func in library.
	f, _ := os.OpenFile("out.png", os.O_WRONLY|os.O_CREATE, 0600)
	defer f.Close()
	png.Encode(f, img)

}

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

	derr := math.Abs(float64(dy)) * 2
	err := 0

	y := y0

	for x := x0; x <= x1; x++ {
		if steep {
			img.Set(y, x, color)
		} else {
			img.Set(x, y, color)
		}

		err += int(derr)
		if err > dx {
			if y1 > y0 {
				y++
			} else {
				y--
			}
			err -= dx * 2
		}
	}
	return img
}

func createTriangle(t0, t1, t2 library.Vector3D, img image.RGBA, color color.RGBA, zBuffer [width*height]int) image.RGBA {
	if t0.Y == t1.Y && t0.Y == t2.Y { return img }

	if t0.Y > t1.Y {
		t0, t1 = t1, t0
	}

	if t0.Y > t2.Y {
		t0, t2 = t2, t0
	}

	if t1.Y > t2.Y {
		t1, t2 = t2, t1
	}

	totalHeight := t2.Y - t0.Y

	for i := 0.0; i < totalHeight; i++ {
		secondHalf := i > t1.Y - t0.Y || t1.Y == t0.Y

		segmentHeight := 0.0

		if secondHalf {
			segmentHeight = t2.Y - t1.Y
		} else {
			segmentHeight = t1.Y - t0.Y
		}

		alpha := i / totalHeight

		beta := 0.0

		if secondHalf {
			beta = (i - (t1.Y - t0.Y)) / segmentHeight
		} else {
			beta = i / segmentHeight
		}

		firstStepA := t2.Subtract(t0)
		secondStepA := firstStepA.MultiplyScalar(alpha)
		A := t0.Add(secondStepA)

		B := library.Vector3D{0,0, 0}

		if secondHalf {
			subVector := t2.Subtract(t1)
			multiplyVector := subVector.MultiplyScalar(beta)
			B = t1.Add(multiplyVector)
		} else {
			subVector := t1.Subtract(t0)
			multiplyVector := subVector.MultiplyScalar(beta)
			B = t0.Add(multiplyVector)
		}

		if A.X > B.X {
			A, B = B, A
		}

		for j := A.X; j <=B.X; j++ {
			var phi = 0.0

			if B.X == A.X {
				phi = 1.0
			} else {
				phi = (j - A.X) / (B.X - A.X)
			}
			subVector := B.Subtract(A)
			multiplyVector := subVector.MultiplyScalar(phi)
			P := A.Add(multiplyVector)

			P.X = j
			P.Y = t0.Y + i

			idx := int(j + P.Y * width)

			if idx <= width*height {
				if zBuffer[idx] < int(P.Z) {
					zBuffer[idx] = int(P.Z)
					img.Set(int(P.X), int(P.Y), color)
				}
			}

			img.Set(int(P.X), int(P.Y), color)

		}
	}
	return img
}
