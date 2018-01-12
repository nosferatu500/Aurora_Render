package _type

import (
	"github.com/nosferatu500/go-vector"
	"github.com/nosferatu500/go-quaternion"
)

var object3DId = 0

type Object3D struct {
	Id int

	Uuid string

	Name string
	Type string

	Parent struct{}
	Children []struct{}

	Up go_vector.Vector3D

	Position go_vector.Vector3D
	Rotation Euler
	Quaternion go_quaternion.Quaternion
	Scale go_vector.Vector3D
}
