package _type

import (
	"github.com/nosferatu500/go-vector"
	"github.com/nosferatu500/go-matrix"
	"AuroraRender/library/utils"
	"image/color"
	"image"
	"fmt"
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
	index := (x + y * bmp.PixelWidth) * 4;
	d.backBuffer[index] = (byte)(color.Blue * 255);
	d.backBuffer[index + 1] = (byte)(color.Green * 255);
	d.backBuffer[index + 2] = (byte)(color.Red * 255);
	d.backBuffer[index + 3] = (byte)(color.Alpha * 255);
}

func Project(coord go_vector.Vector3D, transMat go_matrix.Matrix, bmp *WriteableBitmap) go_vector.Vector2D {
	point := coord.TransformCoordinate(transMat)
fmt.Println(point)
	x := point.X * float64(bmp.PixelWidth) + float64(bmp.PixelWidth) / 2
	y := -point.Y * float64(bmp.PixelHeight) + float64(bmp.PixelHeight) / 2

	return go_vector.Vector2D{x, y }
}

func Render(camera Camera, meshes []Mesh, bmp *WriteableBitmap, img image.RGBA) image.RGBA {
	viewMatrix := utils.LookAtLH(camera.Position, camera.Target, go_vector.Vector3D{0, 1, 0})
	aspect := bmp.PixelWidth / bmp.PixelHeight
	projectionMatrix := go_matrix.PerspectiveFovRH(0.78, float64(aspect), 0.01, 1.0)

	for _, mesh := range meshes {
		worldMatrix := go_matrix.RotationYawPitchRoll(mesh.Rotation.Y, mesh.Rotation.X, mesh.Rotation.Z).Multiply(utils.Translation(mesh.Position))
		transformMatrix := worldMatrix.Multiply(viewMatrix).Multiply(projectionMatrix)



		for _, vertex := range mesh.Vertices {
			fmt.Println(vertex)
			point := Project(vertex, transformMatrix, bmp)
			fmt.Println(point)
			img.Set(int(point.X), int(point.Y), color.RGBA{255,255,0,255})
		}
	}
	return img
}