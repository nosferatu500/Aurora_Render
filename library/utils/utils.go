package utils

import (
	"sync"
	"sync/atomic"
	"runtime"
	"strconv"
	"github.com/nosferatu500/go-vector"
	"github.com/nosferatu500/go-matrix"
)

// parallel starts parallel image processing based on the current GOMAXPROCS value.
func Parallel(dataSize int, fn func(partStart, partEnd int)) {
	numGoroutines := 1
	partSize := dataSize

	numProcs := runtime.GOMAXPROCS(0)
	if numProcs > 1 {
		numGoroutines = numProcs
		partSize = dataSize / (numGoroutines * 10)
		if partSize < 1 {
			partSize = 1
		}
	}

	if numGoroutines == 1 {
		fn(0, dataSize)
	} else {
		var wg sync.WaitGroup
		wg.Add(numGoroutines)
		idx := uint64(0)

		for p := 0; p < numGoroutines; p++ {
			go func() {
				defer wg.Done()
				for {
					partStart := int(atomic.AddUint64(&idx, uint64(partSize))) - partSize
					if partStart >= dataSize {
						break
					}
					partEnd := partStart + partSize
					if partEnd > dataSize {
						partEnd = dataSize
					}
					fn(partStart, partEnd)
				}
			}()
		}

		wg.Wait()
	}
}

func ParseVertex(key string, items []string) go_vector.Vector3D {
	if key != "v" && key != "vt" && key != "vn" && key != "f" { return go_vector.Vector3D{} }

	var Z float64

	if key == "vt" {
		Z = 0.0
	} else {
		Z, _ = strconv.ParseFloat(items[2], 64)
	}

	X, _ := strconv.ParseFloat(items[0], 64)
	Y, _ := strconv.ParseFloat(items[1], 64)

	return go_vector.Vector3D{X, Y, Z}
}

func ParseFace(value string) int {
	parsed, _ := strconv.ParseInt(value, 0, 0)
	n := int(parsed)
	return n
}

func LookAtLH(eye, target, up go_vector.Vector3D) go_matrix.Matrix {
	result1 := target.Subtract(eye).Normalize()
	result2 := up.Cross(result1).Normalize()
	result3 := result1.Cross(result2)

	result := go_matrix.Identity()
	result.M11 = result2.X
	result.M21 = result2.Y
	result.M31 = result2.Z
	result.M12 = result3.X
	result.M22 = result3.Y
	result.M32 = result3.Z
	result.M13 = result1.X
	result.M23 = result1.Y
	result.M33 = result1.Z

	result.M41 = result2.Dot(eye)
	result.M42 = result3.Dot(eye)
	result.M43 = result1.Dot(eye)

	result.M41 = -result.M41
	result.M42 = -result.M42
	result.M43 = -result.M43

	return result
}

func Translation(v go_vector.Vector3D) go_matrix.Matrix {
	result := go_matrix.Identity()
	result.M41 = v.X
	result.M42 = v.Y
	result.M43 = v.Z
	return result
}

