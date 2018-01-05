package _type

import (
	"AuroraRender/library/type/basic"
	"github.com/nosferatu500/go-vector"
)

type Triangle struct {
	V1, V2, V3 basic.Vertex
}

func CreateTriangle(v1, v2, v3 basic.Vertex) *Triangle {
	t := Triangle{v1, v2, v3}
	t.RecoveryNormals()
	return &t
}

func (t *Triangle) GetNormalizeVector3D() go_vector.Vector3D {
	edge1 := t.V2.Position.Subtract(t.V1.Position)
	edge2 := t.V3.Position.Subtract(t.V1.Position)

	return edge1.Cross(edge2).Normalize()
}

func (t *Triangle) RecoveryNormals() {
	normalizeVector := t.GetNormalizeVector3D()
	empty := go_vector.Vector3D{}
	if t.V1.Normal == empty { t.V1.Normal = normalizeVector	}
	if t.V2.Normal == empty { t.V2.Normal = normalizeVector	}
	if t.V3.Normal == empty { t.V3.Normal = normalizeVector	}
}
