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

	var lightDir = library.Vector3D{0,0, -1}

	model := library.CreateModel("./obj/african_head.obj")

	count := library.GetFaceCount(model)

	// Fix bug with out of range for faces collection.
	for i := 0; i < count - 20; i++ {
		face := model.Faces[i]

		var screenCoords [3]library.Vector3D
		var worldCoords [3]library.Vector3D

		for j := 0; j < 3; j++ {
			v := model.Verts[face[j]]

			screenCoords[j] = library.Vector3D{X: (v.X + 1) * width / 2, Y: (v.Y + 1) * height / 2, Z: (v.Z + 1) * depth / 2}
			worldCoords[j] = v
		}

		firstVector := worldCoords[2].Subtract(worldCoords[0])
		secondVector := worldCoords[1].Subtract(worldCoords[0])

		n := firstVector.CrossProduct(secondVector)
		n = n.Normalize()

		intensity := n.Multiply(lightDir)

		newImage = createLine(screenCoords[0], screenCoords[1], image.RGBA{Pix: img.Pix, Stride: img.Stride, Rect: img.Rect}, white)
		if intensity.Z > 0 {
		//	newImage = createTriangle(screenCoords[0], screenCoords[1], screenCoords[2], image.RGBA{Pix: img.Pix, Stride: img.Stride, Rect: img.Rect}, color.RGBA{uint8(intensity.Z) * 255, uint8(intensity.Z) * 255, uint8(intensity.Z) * 255, 255}, zBuffer)
		}
	}

	img = library.FlipByVertically(image.RGBA{Pix: newImage.Pix, Stride: newImage.Stride, Rect: newImage.Rect})

	//TODO: Make as separate func in library.
	f, _ := os.OpenFile("out.png", os.O_WRONLY|os.O_CREATE, 0600)
	defer f.Close()
	png.Encode(f, img)

}

// Return value only for dev test.
func createLine(p0, p1 library.Vector3D, img image.RGBA, color color.RGBA) image.RGBA {
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
