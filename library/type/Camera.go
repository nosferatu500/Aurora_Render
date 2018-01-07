package _type

import "github.com/nosferatu500/go-vector"

type Camera struct {
	Position go_vector.Vector3D
	Target go_vector.Vector3D
}

func CreateCamera(position, target go_vector.Vector3D) Camera {
	return Camera{Position: position, Target: target}
}
