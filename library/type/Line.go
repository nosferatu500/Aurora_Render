package _type

import (
	"AuroraRender/library/type/basic"
)

type Line struct {
	V1, V2 basic.Vertex
}

func CreateLine(v1, v2 basic.Vertex) *Line {
	return &Line{v1, v2}
}
