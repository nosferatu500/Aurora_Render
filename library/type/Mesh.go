package _type

type Mesh struct {
	Triangles []*Triangle
	Lines     []*Line
	box       *Box
}

func CreateMesh(triangles []*Triangle) *Mesh {
	return &Mesh{triangles, nil, nil}
}
