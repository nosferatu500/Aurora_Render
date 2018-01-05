package utils

import (
	"image"
)

func FlipByVertically(img image.RGBA) *image.RGBA {
	src := img
	srcW := src.Bounds().Max.X
	srcH := src.Bounds().Max.Y
	dstW := srcW
	dstH := srcH
	dst := image.NewRGBA(image.Rect(0, 0, dstW, dstH))

	Parallel(dstH, func(partStart, partEnd int) {

		for dstY := partStart; dstY < partEnd; dstY++ {
			for dstX := 0; dstX < dstW; dstX++ {
				srcX := dstX
				srcY := dstH - dstY - 1

				srcOff := srcY*src.Stride + srcX*4
				dstOff := dstY*dst.Stride + dstX*4

				copy(dst.Pix[dstOff:dstOff+4], src.Pix[srcOff:srcOff+4])
			}
		}

	})

	return dst
}
