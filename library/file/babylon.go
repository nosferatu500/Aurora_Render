package file

import (
	"AuroraRender/library/type"
	"io/ioutil"
	"fmt"
	"encoding/json"
	"github.com/nosferatu500/go-vector"
	"AuroraRender/library/type/basic"
)

type Babylon struct {
	Meshes []struct{
		Name string
		Vertices []float64
		Indices []int
		UvCount int
		Position []float64
	}
}

func LoadBabylon(path string) ([]*_type.Mesh, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var babylon Babylon

	fmt.Println("data",data)

	json.Unmarshal(data, &babylon)

	fmt.Println("babylon",babylon)

	var meshes []*_type.Mesh

	for meshIndex := 0; meshIndex < len(babylon.Meshes); meshIndex++ {
		verticesArray := babylon.Meshes[meshIndex].Vertices

		indicesArray := babylon.Meshes[meshIndex].Indices

		uvCount := babylon.Meshes[meshIndex].UvCount

		verticesStep := 1

		switch uvCount {
		case 0:
		verticesStep = 6
			break
		case 1:
		verticesStep = 8
			break
		case 2:
		verticesStep = 10
			break
		}

		verticesCount := len(verticesArray) / verticesStep

		facesCount := len(indicesArray) / 3

		var vertices []go_vector.Vector3D
		var faces []basic.Face

		for index := 0; index < verticesCount; index++ {
			x := verticesArray[index * verticesStep]
			y := verticesArray[index * verticesStep + 1]
			z := verticesArray[index * verticesStep + 2]
			vertices = append(vertices, go_vector.Vector3D{x, y, z})
		}

		for index := 0; index < facesCount; index++ {
			a := indicesArray[index * 3]
			b := indicesArray[index * 3 + 1]
			c := indicesArray[index * 3 + 2]
			faces = append(faces, basic.Face{ A: a, B: b, C: c })
		}

		mesh := _type.CreateMesh(babylon.Meshes[meshIndex].Name, vertices, faces)

		position := babylon.Meshes[meshIndex].Position;
		mesh.Position = go_vector.Vector3D{position[0], position[1], position[2]}
		meshes = append(meshes, mesh)
	}

	return meshes, err
}
