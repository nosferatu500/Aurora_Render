package file

import (
	"strings"
	"os"
	"bufio"
	"strconv"
	"github.com/nosferatu500/go-vector"
	"AuroraRender/library/type"
	"AuroraRender/library/utils"
	"fmt"
)

func parseIndex(value string, length int) int {
	parsed, _ := strconv.ParseInt(value, 0, 0)
	n := int(parsed)
	if n < 0 { n += length }
	return n
}

func LoadOBJ(path string) (*_type.Mesh, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	vs := make([]go_vector.Vector3D, 1, 1024)  // 1-based indexing
	vts := make([]go_vector.Vector3D, 1, 1024) // 1-based indexing
	vns := make([]go_vector.Vector3D, 1, 1024) // 1-based indexing
	var triangles []*_type.Triangle
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)

		if len(fields) == 0 { continue }

		key := fields[0]
		args := fields[1:]

		v := utils.ParseVertex(key, args)

		switch key {
		case "v":	vs = append(vs, v)
		case "vt":	vts = append(vts, v)
		case "vn":	vns = append(vns, v)
		case "f":
			fvs := make([]int, len(args))
			fvts := make([]int, len(args))
			fvns := make([]int, len(args))
			for i, arg := range args {
				vertex := strings.Split(arg+"//", "/")

				fvs[i] = parseIndex(vertex[0], len(vs))
				fvts[i] = parseIndex(vertex[1], len(vts))
				fvns[i] = parseIndex(vertex[2], len(vns))

				fmt.Println(i, arg, vertex, vs[fvs[0]], len(vs))
			}
			for i := 1; i < len(fvs)-1; i++ {
				i1, i2, i3 := 0, i, i+1
				t := _type.Triangle{}
				t.V1.Position = vs[fvs[i1]]
				t.V2.Position = vs[fvs[i2]]
				t.V3.Position = vs[fvs[i3]]
				t.V1.Normal = vns[fvns[i1]]
				t.V2.Normal = vns[fvns[i2]]
				t.V3.Normal = vns[fvns[i3]]
				t.V1.Texture = vts[fvts[i1]]
				t.V2.Texture = vts[fvts[i2]]
				t.V3.Texture = vts[fvts[i3]]
				t.RecoveryNormals()
				triangles = append(triangles, &t)
			}
		}
	}
	return _type.CreateMesh(triangles), scanner.Err()
}
