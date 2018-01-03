package library

import (
	"fmt"
	"math"
)

type Vector3D struct {
	X float64
	Y float64
	Z float64
}

func (v Vector3D) Add(v2 Vector3D) Vector3D {
	return Vector3D{v.X + v2.X, v.Y + v2.Y, v.Z + v2.Z}
}

func (v Vector3D) Subtract(v2 Vector3D) Vector3D {
	return Vector3D{v.X - v2.X, v.Y - v2.Y, v.Z - v2.Z}
}

func (v Vector3D) CrossProduct(v2 Vector3D) Vector3D  {
	x := (v.Y * v2.Z) - (v.Z * v2.Y)
	y := (v.Z * v2.X) - (v.X * v2.Z)
	z := (v.X * v2.Y) - (v.Y * v2.X)
	return Vector3D{x, y, z}
}

func (v Vector3D) Normalize() Vector3D  {
	length := math.Sqrt(v.X * v.X + v.Y * v.Y + v.Z * v.Z)
	x := v.X / length
	y := v.Y / length
	z := v.Z / length
	return Vector3D{x, y, z}
}

func (v Vector3D) Multiply(v2 Vector3D) Vector3D {
	return Vector3D{v.X * v2.X, v.Y * v2.Y, v.Z * v2.Z}
}

func (v Vector3D) String() string {
	return fmt.Sprintf("%v:%v:%v", v.X, v.Y, v.Z)
}
