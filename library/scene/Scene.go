package scene

import "AuroraRender/library/type"

type Scene struct {
	_type.Object3D

	background struct{}
	fog struct{}
	overrideMaterial struct{}

	autoUpdate bool
}

func CreateScene() Scene {
	return Scene{Object3D: _type.Object3D{Type: "Scene"}, autoUpdate: true}
}
