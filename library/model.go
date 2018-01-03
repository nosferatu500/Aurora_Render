package library

import (
	"os"
	"bufio"
	"fmt"
	"log"
	"regexp"
	"strings"
	"strconv"
)

type Model struct {
	Verts []Vector3D
	Faces []Face
}

func CreateModel(path string) Model {
	var model Model

	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var faces []Face
	var vertexes []Vector3D

	for scanner.Scan() {
		line := scanner.Text()

		// Collect faces from file.
		if strings.Contains(line, "f ") {
			re := regexp.MustCompile("[0-9]+")
			coords := re.FindAllString(line, -1)

			var newFace []int

			for i := 0; i < len(coords); i++ {
				if len(coords) == 9 {
					if i == 0 || i == 3 || i == 6 {
						coord, err := strconv.ParseFloat(coords[i], 10)

						if err != nil {
							fmt.Println("Error in creation model process", err)
						}

						newFace = append(newFace, int(coord))
					}
				}
			}

			faces = append(faces, newFace)
		}

		// Collect vertexes from file.
		if strings.Contains(line, "v ") {
			re := regexp.MustCompile("[+-]?([0-9]*[.])?[0-9]+")
			coords := re.FindAllString(line, -1)

			var newVert Vector3D

			for i := 0; i < len(coords); i++ {
				if len(coords) == 3 {
					coord, err := strconv.ParseFloat(coords[i], 10)

					if err != nil {
						fmt.Println("Error in creation model process", err)
					}

					switch i {
					case 0:
						newVert.X = coord
						break
					case 1:
						newVert.Y = coord
						break
					case 2:
						newVert.Z = coord
						break
					default:
						fmt.Println("Collection of vertixes is incorrect.")
					}
				}
			}

			vertexes = append(vertexes, newVert)
		}
	}

	model = Model{vertexes, faces}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return model
}

func GetVertexCount(model Model) int {
	return len(model.Verts)
}

func GetFaceCount(model Model) int {
	return len(model.Faces)
}
