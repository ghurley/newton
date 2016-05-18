package newton

import (
	"image"
	"image/color"
	"math/cmplx"
)

// .866025403 = math.Sqrt(3.0)/2.0
const (
	root1 = -0.5 + .866025403i
	root2 = -0.5 - .866025403i
	root3 = 1.0 + 0.0i
)

func newton64(x complex64) complex64 {
	// The cmplx methods only operate on complex128s so we although we're
	// interetesting in looking at the degredation caused by using complex64,
	// we still have to cast up to the 128 value. We'll still be operating a
	// defecit but there's a better way to do this
	x128 := complex128(x)
	return complex64(x128 - ((cmplx.Pow(x128, 3) - 1) / (3.0 * cmplx.Pow(x128, 2))))
}

// DrawNewton64 iterates over complex points in a plane defined by two corner
// values and divided into a number of horizontal and vertical pixels. For
// each point, call getColor() for the given point. Return an RGBA.
func DrawNewton64(numRealValues, numImagValues int, start, end complex64) *image.RGBA {
	realStepValue := (real(end) - real(start)) / float32(numRealValues)
	imagStepValue := (imag(end) - imag(start)) / float32(numImagValues)

	img := image.NewRGBA(image.Rect(0, 0, numRealValues, numImagValues))

	for x := 0; x < numRealValues; x++ {
		rc := real(start) + realStepValue*float32(x)
		for y := 0; y < numImagValues; y++ {
			ic := imag(start) + imagStepValue*float32(y)
			color := getColor64(complex(rc, ic))
			img.Set(x, y, color)
		}
	}
	return img
}

// GetColor returns a color for a pixel/point represented by the input p.
func getColor64(p complex64) color.Color {
	const epsilon = 0.000001
	const contrast = 12

	pv := complex128(newton64(p))
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
		pv = complex128(newton64(complex64(pv)))
	}

	return color.RGBA{0x00, 0x00, 0x00, 0xff}
}
