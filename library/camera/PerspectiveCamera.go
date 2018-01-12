package camera

type PerspectiveCamera struct {
	Camera

	Fov int
	Zoom int

	Near float64
	Far int
	Focus int

	Aspect int

	View struct{}

	FilmGauge int
	FilmOffset int
}

func CreatePerspectiveCamera(fov, aspect, far int, near float64) PerspectiveCamera {
	return PerspectiveCamera{CreateCamera("PerspectiveCamera"), fov, 1, near, far, 10, aspect, struct{}{}, 35, 0}
}
