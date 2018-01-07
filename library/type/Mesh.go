package _type

import "github.com/nosferatu500/go-vector"

type Mesh struct {
	Name string
	Vertices []go_vector.Vector3D
	Position go_vector.Vector3D
	Rotation go_vector.Vector3D
}

func CreateMesh(name string, vertices []go_vector.Vector3D) *Mesh {
	return &Mesh{name, vertices, go_vector.Vector3D{0,0,0}, go_vector.Vector3D{0,0,0}}
}
