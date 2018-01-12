package camera

import (
	"github.com/nosferatu500/go-matrix"
	"AuroraRender/library/type"
)

type Camera struct {
	_type.Object3D
	_type string
	matrixWorldInverse go_matrix.Matrix
	projectionMatrix go_matrix.Matrix
}

func CreateCamera(_type string) Camera {
	return Camera{_type.Object3D{Type: "Camera"}, _type, go_matrix.Identity(), go_matrix.Identity()}
}
