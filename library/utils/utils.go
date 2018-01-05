package utils

import (
	"sync"
	"sync/atomic"
	"runtime"
	"strconv"
	"github.com/nosferatu500/go-vector"
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
