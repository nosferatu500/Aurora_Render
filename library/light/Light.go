package light

import (
	"AuroraRender/library/type"
	"AuroraRender/library/type/basic"
)

type Light struct {
	_type.Object3D
	color basic.Color
	intensity int
	receiveShadow bool
}

func CreateLight(color basic.Color, intensity int) Light {
	return Light{Object3D: _type.Object3D{Type: "Light"}, color: basic.Color{color.R, color.G, color.B, color.A}, intensity: intensity}
}