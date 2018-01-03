package library

import "fmt"

type Vector3D struct {
	X float64
	Y float64
	Z float64
}

func (v Vector3D) String() string {
	return fmt.Sprintf("%v:%v:%v", v.X, v.Y, v.Z)
}
