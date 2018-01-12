package light

import "AuroraRender/library/type/basic"

type AmbientLight struct {
	Light
	castShadow bool
}

func CreateAmbientLight(color basic.Color, intensity int) AmbientLight {
	return AmbientLight{CreateLight(color, intensity), false}
}