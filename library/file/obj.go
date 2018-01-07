package file

import (
	"strings"
	"os"
	"bufio"
	"github.com/nosferatu500/go-vector"
	"AuroraRender/library/type"
	"AuroraRender/library/type/basic"
	"strconv"
)

//Fix empty faces
func LoadOBJ(path string) (*_type.Mesh, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var vertices []go_vector.Vector3D
	var faces []basic.Face

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		switch {
		case strings.HasPrefix(line, "f "):
			face := parseFace(line)
			faces = append(faces, face)
		case strings.HasPrefix(line, "v "):
			vertex := parseVertex(line)
			vertices = append(vertices, vertex)
		}
	}
	return _type.CreateMesh("test", vertices, faces), err
}

func parseFace(line string) basic.Face {
	parts := strings.Split(line, " ")[1:] // Skip the initial "f".
	var indices []int

	for _, part := range parts {
		idx := strings.Split(part, "/")[0]
		ind, _ := strconv.Atoi(idx)
		indices = append(indices, ind)
	}

	return basic.Face{indices[0],indices[1],indices[2]}
}

func parseVertex(line string) go_vector.Vector3D {
	parts := strings.Split(line, " ")[1:] // Skip the initial "v".

	X, _ := strconv.ParseFloat(parts[0], 64)
	Y, _ := strconv.ParseFloat(parts[1], 64)
	Z, _ := strconv.ParseFloat(parts[2], 64)

	return go_vector.Vector3D{X, Y, Z}
}
