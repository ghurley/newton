package newton

import (
	"image"
	"image/color"
	"math/cmplx"
)

// .866025403 = math.Sqrt(3.0)/2.0
const (
	root1_128 = -0.5 + .866025403i
	root2_128 = -0.5 - .866025403i
	root3_128 = 1.0 + 0.0i
)

func newton128(x complex128) complex128 {
	return x - ((cmplx.Pow(x, 3) - 1) / (3.0 * cmplx.Pow(x, 2)))
}

// DrawNewton128 iterates over complex points in a plane defined by two corner
// values and divided into a number of horizontal and vertical bounds. For
// each point, call GetColor() for the given point. Return an RGBA.
func DrawNewton128(numRealValues, numImagValues int, start, end complex128) *image.RGBA {
	realStepValue := (real(end) - real(start)) / float64(numRealValues)
	imagStepValue := (imag(end) - imag(start)) / float64(numImagValues)

	img := image.NewRGBA(image.Rect(0, 0, numRealValues, numImagValues))

	for x := 0; x < numRealValues; x++ {
		rc := real(start) + realStepValue*float64(x)
		for y := 0; y < numImagValues; y++ {
			ic := imag(start) + imagStepValue*float64(y)
			color := getColor128(complex(rc, ic))
			img.Set(x, y, color)
		}
	}
	return img
}

// GetColor returns a color for a pixel/point represented by the input c.
func getColor128(p complex128) color.Color {
	const epsilon = 0.000001
	const contrast = 12

	pv := newton128(p)
	for iterations := uint8(0); iterations < 20; iterations++ {
		activeColorLevel := 255 - contrast*iterations
		switch {
		case cmplx.Abs(pv-root1) < epsilon:
			return color.RGBA{0x00, 0x00, activeColorLevel, 0xff}
		case cmplx.Abs(pv-root2) < epsilon:
			return color.RGBA{0x00, activeColorLevel, 0x00, 0xff}
		case cmplx.Abs(pv-root3) < epsilon:
			return color.RGBA{activeColorLevel, 0x00, 0x00, 0xff}
		}
		pv = newton128(pv)
	}

	return color.RGBA{0x00, 0x00, 0x00, 0xff}
}
