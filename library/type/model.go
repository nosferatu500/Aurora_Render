package _type

import (
	"os"
	"bufio"
	"log"
	"strings"
	"strconv"
	"github.com/nosferatu500/go-vector"
	"AuroraRender/library/type/basic"
)

type Model struct {
	Verts []go_vector.Vector3D
	Faces []basic.Face
}

func CreateModel(path string) Model {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var vertexes []go_vector.Vector3D
	var faces []basic.Face

	for scanner.Scan() {
		line := scanner.Text()

		switch {
		case strings.HasPrefix(line, "f "):
			face := parseFace(line)
			faces = append(faces, face)
		case strings.HasPrefix(line, "v "):
			vertex := parseVertex(line)
			vertexes = append(vertexes, vertex)
		}

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return Model{vertexes, faces}
}

func GetVertexCount(model Model) int {
	return len(model.Verts)
}

func GetFaceCount(model Model) int {
	return len(model.Faces)
}

func parseFace(line string) basic.Face {
	parts := strings.Split(line, " ")[1:] // Skip the initial "f".
	indices := make([]int, len(parts))

	for i, part := range parts {
		idx := strings.Split(part, "/")[0]
		indices[i], _ = strconv.Atoi(idx)
	}

	return basic.Face{0,0,0}
}

func parseVertex(line string) go_vector.Vector3D {
	parts := strings.Split(line, " ")[1:] // Skip the initial "v".

	X, _ := strconv.ParseFloat(parts[0], 64)
	Y, _ := strconv.ParseFloat(parts[1], 64)
	Z, _ := strconv.ParseFloat(parts[2], 64)

	return go_vector.Vector3D{X, Y, Z}
}
