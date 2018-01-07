package _type

import (
	"github.com/nosferatu500/go-vector"
	"AuroraRender/library/type/basic"
)

type Mesh struct {
	Name string
	Vertices []go_vector.Vector3D
	Faces []basic.Face
	Position go_vector.Vector3D
	Rotation go_vector.Vector3D
}

func CreateMesh(name string, vertices []go_vector.Vector3D, faces []basic.Face) *Mesh {
	return &Mesh{name, vertices, faces, go_vector.Vector3D{0,0,0}, go_vector.Vector3D{0,0,0}}
}
