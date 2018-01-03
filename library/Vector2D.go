package library

import "fmt"

type Vector2D struct {
	X float64
	Y float64
}

func (v Vector2D) Add(v2 Vector2D) Vector2D {
	return Vector2D{v.X + v2.X, v.Y + v2.Y}
}

func (v Vector2D) Subtract(v2 Vector2D) Vector2D {
	return Vector2D{v.X - v2.X, v.Y - v2.Y}
}

func (v Vector2D) MultiplyScalar(s float64) Vector2D {
	return Vector2D{v.X * s, v.Y * s}
}

func (v Vector2D) String() string {
	return fmt.Sprintf("%v:%v", v.X, v.Y)
}
