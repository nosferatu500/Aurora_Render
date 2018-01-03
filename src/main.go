package main

import (
	"image"
	"image/color"
	"os"
	"image/png"
	"math"
	"AuroraRender/library"
)

const (
	width = 1000
	height = 1000
)

//TODO: Make as const in future release.
var white = color.RGBA{255, 255, 255, 255}
var red = color.RGBA{255, 0, 0, 255}
var blue = color.RGBA{0, 0, 255, 255}

func main() {
	img := image.NewRGBA(image.Rect(0, 0, width, height))

	var newImage image.RGBA

	var lightDir = library.Vector3D{0,0, -1}

	model := library.CreateModel("./obj/african_head.obj")

	count := library.GetFaceCount(model)

	// Fix bug with out of range for faces collection.
	for i := 0; i < count - 20; i++ {
		face := model.Faces[i]

		var screenCoords [3]library.Vector2D
		var worldCoords [3]library.Vector3D

		for j := 0; j < 3; j++ {
			v := model.Verts[face[j]]

			screenCoords[j] = library.Vector2D{X: (v.X + 1) * width / 2, Y: (v.Y + 1) * height / 2}
			worldCoords[j] = v
		}

		n := (worldCoords[2].Subtract(worldCoords[0])).CrossProduct(worldCoords[1].Subtract(worldCoords[0]))
		n = n.Normalize()

		intensity := n.Multiply(lightDir)

		newImage = createLine(screenCoords[0], screenCoords[1], image.RGBA{Pix: img.Pix, Stride: img.Stride, Rect: img.Rect}, white)
		if intensity.Z > 0 {
			newImage = createTriangle(screenCoords[0], screenCoords[1], screenCoords[2], image.RGBA{Pix: newImage.Pix, Stride: newImage.Stride, Rect: newImage.Rect}, color.RGBA{uint8(intensity.Z) * 255, uint8(intensity.Z) * 255, uint8(intensity.Z) * 255, 255})
		}
	}

	img = library.FlipByVertically(image.RGBA{Pix: newImage.Pix, Stride: newImage.Stride, Rect: newImage.Rect})

	//TODO: Make as separate func in library.
	f, _ := os.OpenFile("out.png", os.O_WRONLY|os.O_CREATE, 0600)
	defer f.Close()
	png.Encode(f, img)

}

// Return value only for dev test.
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
		t := (x - p0.X) / (p1.X - p0.X)

		y := p0.Y * (1 - t) + p1.Y * t

		if steep {
			img.Set(int(y), int(x), color)
		} else {
			img.Set(int(x), int(y), color)
		}
	}
	return img
}

func createTriangle(t0, t1, t2 library.Vector2D, img image.RGBA, color color.RGBA) image.RGBA {
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

		A := t0.Add(t2.Subtract(t0)).MultiplyScalar(alpha)

		B := library.Vector2D{0,0}

		if secondHalf {
			B = t1.Add(t2.Subtract(t1)).MultiplyScalar(beta)
		} else {
			B = t0.Add(t1.Subtract(t0)).MultiplyScalar(beta)
		}

		if A.X > B.X {
			A, B = B, A
		}

		for j := A.X; j <=B.X; j++ {
			img.Set(int(j), int(t0.Y + i), color)
		}
	}
	return img
}
