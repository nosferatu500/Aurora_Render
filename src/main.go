package main

import (
	"image"
	"image/color"
	"os"
	"image/png"
	"math"
	"github.com/nosferatu500/go-vector"
	"fmt"
	//"strconv"
	"AuroraRender/library/type"
	"AuroraRender/library/utils"
	//"AuroraRender/library/file"
	"AuroraRender/library/type/basic"
)

const (
	width = 1920
	height = 1920
	depth = 255
	SSAA = 4
)

//TODO: Make as const in future release.
var white = color.RGBA{255, 255, 255, 255}
var red = color.RGBA{255, 0, 0, 255}
var blue = color.RGBA{0, 0, 255, 255}

func main() {
	img := image.NewRGBA(image.Rect(0, 0, width, height))

	bmp := _type.WriteableBitmap{width, height, nil}

	buffer := make([]byte, bmp.PixelHeight * bmp.PixelWidth * 4)

	device := _type.CreateDevice(buffer, bmp)

	camera := _type.CreateCamera(go_vector.Vector3D{0,0,10}, go_vector.Vector3D{0,0,0})

	var vertices []go_vector.Vector3D

	vertices = append(vertices, go_vector.Vector3D{-1, 1, 1})
	vertices = append(vertices, go_vector.Vector3D{1, 1, 1})
	vertices = append(vertices, go_vector.Vector3D{-1, -1, 1})
	vertices = append(vertices, go_vector.Vector3D{-1, -1, -1})
	vertices = append(vertices, go_vector.Vector3D{-1, 1, -1})
	vertices = append(vertices, go_vector.Vector3D{1, 1, -1})
	vertices = append(vertices, go_vector.Vector3D{1, -1, 1})
	vertices = append(vertices, go_vector.Vector3D{1, -1, -1})

	var faces []basic.Face

	faces = append(faces, basic.Face{A:0, B: 1, C: 2})
	faces = append(faces, basic.Face{A:1, B: 2, C: 3})
	faces = append(faces, basic.Face{A:1, B: 3, C: 6})
	faces = append(faces, basic.Face{A:1, B: 5, C: 6})
	faces = append(faces, basic.Face{A:0, B: 1, C: 4})
	faces = append(faces, basic.Face{A:1, B: 4, C: 5})

	faces = append(faces, basic.Face{A:2, B: 3, C: 7})
	faces = append(faces, basic.Face{A:3, B: 6, C: 7})
	faces = append(faces, basic.Face{A:0, B: 2, C: 7})
	faces = append(faces, basic.Face{A:0, B: 4, C: 7})
	faces = append(faces, basic.Face{A:4, B: 5, C: 6})
	faces = append(faces, basic.Face{A:4, B: 6, C: 7})

	newMesh := _type.CreateMesh("Cube", vertices, faces)

	device.Clear(0,0,0,255)
	newMesh.Rotation = go_vector.Vector3D{newMesh.Rotation.X + 0.01, newMesh.Rotation.Y + 0.01, newMesh.Rotation.Z}

	meshes := make([]_type.Mesh, 0)

	meshes = append(meshes, *newMesh)
	var newImage image.RGBA
	newImage = _type.Render(camera, meshes, &bmp, *img)

	bmp.Present(device.GetBackBuffer())




	var zBuffer [width*height]int

	for i:=0; i<width*height; i++ {
		zBuffer[i] = -2147483648
	}

	model := _type.CreateModel("./obj/african_head.obj") // ~2.5K triangles
	// model := _type.CreateModel("./obj/torso2_base_fin.obj")

	//mesh, _ := file.LoadOBJ("./obj/african_head.obj")
	// mesh, _ := file.LoadOBJ("./obj/torso2_base_fin.obj") // ~1M triangles

	//fmt.Println(len(mesh.Triangles), "mesh")

	count := _type.GetFaceCount(model)
	fmt.Println(count, "count")

/*
	for i := 0; i < count - 20; i++ {
		face := model.Faces[i]

		var screenCoords [3]go_vector.Vector3D
		var worldCoords [3]go_vector.Vector3D

		for j := 0; j < 3; j++ {
			v := model.Verts[face[j] - 1]
			screenCoords[j] = go_vector.Vector3D{(v.X + 1) * width / 2, (v.Y + 1) * height / 2, (v.Z + 1) * depth / 2}
			worldCoords[j] = v
		}

		edge1 := worldCoords[2].Subtract(worldCoords[0])
		edge2 := worldCoords[1].Subtract(worldCoords[0])

		n := edge1.Cross(edge2).Normalize()

		intensity := n.Multiply(go_vector.Vector3D{0,0,-1})

		// newImage = createLine(screenCoords[1], screenCoords[2], image.RGBA{Pix: img.Pix, Stride: img.Stride, Rect: img.Rect}, white)

		//r := rand.New(rand.NewSource(int64(i)))
		//randomColor := color.RGBA{uint8(r.Int()), uint8(r.Int()), uint8(r.Int()), 255}
		//newImage = createTriangle(screenCoords[0], screenCoords[1], screenCoords[2],image.RGBA{Pix: img.Pix, Stride: img.Stride, Rect: img.Rect}, randomColor)



		if intensity.Z > 0 {
			newNumberString := strconv.FormatFloat(intensity.Z * 255, 'f', 0, 64)
			newNumber, _ := strconv.ParseInt(newNumberString, 10, 8)

			intensityColor := color.RGBA{uint8(newNumber), uint8(newNumber), uint8(newNumber), 255}
			newImage = createTriangle(screenCoords[0], screenCoords[1], screenCoords[2],image.RGBA{Pix: img.Pix, Stride: img.Stride, Rect: img.Rect}, intensityColor)
		}
	}
*/
/*
	for i := 0; i < len(mesh.Triangles); i++ {
		fmt.Println(i, " of ", len(mesh.Triangles))
		triangle := mesh.Triangles[i]

		var screenCoords [3]go_vector.Vector3D
		var worldCoords [3]go_vector.Vector3D

		for j := 0; j < 3; j++ {
			v := go_vector.Vector3D{}

			switch j {
			case 0: v = triangle.V1.Position
			case 1: v = triangle.V2.Position
			case 2: v = triangle.V3.Position
			}

			screenCoords[j] = go_vector.Vector3D{(v.X + 1) * width / 2, (v.Y + 1) * height / 2, (v.Z + 1) * depth / 2}
			worldCoords[j] = v
		}

		edge1 := worldCoords[2].Subtract(worldCoords[0])
		edge2 := worldCoords[1].Subtract(worldCoords[0])

		n := edge1.Cross(edge2).Normalize()

		intensity := n.Multiply(go_vector.Vector3D{0,0,-0.5})

		// newImage = createLine(screenCoords[1], screenCoords[2], image.RGBA{Pix: img.Pix, Stride: img.Stride, Rect: img.Rect}, white)

		//r := rand.New(rand.NewSource(int64(i)))
		//randomColor := color.RGBA{uint8(r.Int()), uint8(r.Int()), uint8(r.Int()), 255}
		//newImage = createTriangle(screenCoords[0], screenCoords[1], screenCoords[2],image.RGBA{Pix: img.Pix, Stride: img.Stride, Rect: img.Rect}, randomColor)

		if intensity.Z > 0 {
			newNumberString := strconv.FormatFloat(intensity.Z * 255, 'f', 0, 64)
			newNumber, _ := strconv.ParseInt(newNumberString, 10, 8)

			intensityColor := color.RGBA{uint8(newNumber), uint8(newNumber), uint8(newNumber), 255}
			newImage = createTriangle(screenCoords[0], screenCoords[1], screenCoords[2],image.RGBA{Pix: img.Pix, Stride: img.Stride, Rect: img.Rect}, intensityColor)
		}
	}

*/

	img = utils.FlipByVertically(image.RGBA{Pix: newImage.Pix, Stride: newImage.Stride, Rect: newImage.Rect})

	//TODO: Make as separate func in library.
	f, _ := os.OpenFile("out.png", os.O_WRONLY|os.O_CREATE, 0600)
	defer f.Close()
	png.Encode(f, img)

	fmt.Println("Write file complete.")
}

func createLine(p0, p1 go_vector.Vector2D, img image.RGBA, color color.RGBA) image.RGBA {
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

func createTriangle(t0, t1, t2 go_vector.Vector3D, img image.RGBA, color color.RGBA) image.RGBA {
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
