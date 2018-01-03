package library

import "fmt"

type Vector2D struct {
	X float64
	Y float64
}

func (v Vector2D) String() string {
	return fmt.Sprintf("%v:%v", v.X, v.Y)
}
