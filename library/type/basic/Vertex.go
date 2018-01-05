package basic

import "github.com/nosferatu500/go-vector"

type Vertex struct {
	Position go_vector.Vector3D
	Normal   go_vector.Vector3D
	Texture  go_vector.Vector3D
	Color    Color
	// Output   VectorW
	// Vectors  []Vector
	// Colors   []Color
	// Floats   []float64
}
