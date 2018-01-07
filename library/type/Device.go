package _type

import (
	"github.com/nosferatu500/go-vector"
	"github.com/nosferatu500/go-matrix"
	"AuroraRender/library/utils"
	"image/color"
	"image"
	"fmt"
	"math"
)

type Device struct {
	backBuffer []byte // bmp.PixelHeight * bmp.PixelWidth * 4
	bmp WriteableBitmap
}

func (d Device) GetBackBuffer() []byte {
	return d.backBuffer
}

func CreateDevice(backBuffer []byte, bmp WriteableBitmap) Device {
	return Device{backBuffer, bmp}
}

func (d Device) Clear(r, g, b, a byte) {
	for i := 0; i < len(d.backBuffer); i+=4 {
		d.backBuffer[i] = b
		d.backBuffer[i + 1] = g
		d.backBuffer[i + 2] = r
		d.backBuffer[i + 3] = a
	}
}

func (bmp WriteableBitmap) Present(backBuffer []byte) {
	bmp.PixelBuffer = backBuffer
}

func (bmp WriteableBitmap) PutPixel(x, y int, color Color4, d *Device)  {
	index := (x + y * bmp.PixelWidth) * 4
	d.backBuffer[index] = (byte)(color.Blue * 255)
	d.backBuffer[index + 1] = (byte)(color.Green * 255)
	d.backBuffer[index + 2] = (byte)(color.Red * 255)
	d.backBuffer[index + 3] = (byte)(color.Alpha * 255)
}

func Project(coord go_vector.Vector3D, transMat go_matrix.Matrix, bmp *WriteableBitmap) go_vector.Vector2D {
	point := coord.TransformCoordinate(transMat)
fmt.Println(point)
	x := point.X * float64(bmp.PixelWidth) + float64(bmp.PixelWidth) / 2
	y := -point.Y * float64(bmp.PixelHeight) + float64(bmp.PixelHeight) / 2

	return go_vector.Vector2D{x, y }
}

func DrawLine(point0, point1 go_vector.Vector2D, img image.RGBA) image.RGBA {
	dist := point1.Subtract(point0).Length()

	if dist < 2 { return img}

	middlePoint := point1.Subtract(point0).DivideScalar(2).Add(point0)
	img.Set(int(middlePoint.X), int(middlePoint.Y), color.RGBA{255,255,0,255})

	DrawLine(point0, middlePoint, img)
	DrawLine(middlePoint, point1, img)

	return img
}

func DrawBLine(point0, point1 go_vector.Vector2D, img image.RGBA) image.RGBA {
	x0 := int(point0.X)
	y0 := int(point0.Y)

	x1 := int(point1.X)
	y1 := int(point1.Y)

	dx := math.Abs(float64(x1) - float64(x0))
	dy := math.Abs(float64(y1) - float64(y0))

	sx := 0.0
	sy := 0.0

	if x0 < x1  { sx = 1 } else { sx = -1 }
	if y0 < y1  { sy = 1 } else { sy = -1 }

	err := dx - dy

	for true {
		img.Set(int(x0), int(y0), color.RGBA{255,255,0,255})

		if (x0 == x1) && (y0 == y1) { break }

		e2 := 2 * err

		if e2 > -dy { err -= dy; x0 += int(sx) }
		if e2 < dx { err += dx; y0 += int(sy) }
	}
	return img
}

func Render(camera Camera, meshes []*Mesh, bmp *WriteableBitmap, img image.RGBA) image.RGBA {
	viewMatrix := utils.LookAtLH(camera.Position, camera.Target, go_vector.Vector3D{0, 1, 0})
	aspect := bmp.PixelWidth / bmp.PixelHeight
	projectionMatrix := go_matrix.PerspectiveFovRH(0.78, float64(aspect), 0.01, 1.0)

	for _, mesh := range meshes {
		worldMatrix := go_matrix.RotationYawPitchRoll(mesh.Rotation.Y, mesh.Rotation.X, mesh.Rotation.Z).Multiply(utils.Translation(mesh.Position))
		transformMatrix := worldMatrix.Multiply(viewMatrix).Multiply(projectionMatrix)
		fmt.Println(len(mesh.Faces), len(mesh.Vertices))

		for i, face := range mesh.Faces {
			fmt.Println("face.B", face.B)
			fmt.Println(i, len(mesh.Faces))

			vertexA := mesh.Vertices[face.A]
			vertexB := mesh.Vertices[face.B]
			vertexC := mesh.Vertices[face.C]


			pixelA := Project(vertexA, transformMatrix, bmp)
			pixelB := Project(vertexB, transformMatrix, bmp)
			pixelC := Project(vertexC, transformMatrix, bmp)

			img = DrawBLine(pixelA, pixelB, img)
			img = DrawBLine(pixelB, pixelC, img)
			img = DrawBLine(pixelC, pixelA, img)
		}
	}
	return img
}