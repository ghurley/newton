package main

import (
	"fmt"
	"image/png"
	"log"
	"os"
	"runtime/pprof"
	"time"

	"github.com/ghurley/newton"
)

func main() {
	const upperLeft = -1.5 + 1.5i
	const lowerRight = 1.5 - 1.5i

	// Un/comment for profiling.
	f, err := os.Create("prog128.prof")
	if err != nil {
		log.Fatal(err)
	}
	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()
	// End profiling code.

	start := time.Now()

	img := newton.DrawNewton128(640, 480, upperLeft, lowerRight)
	fmt.Println("Calculation time: ", time.Since(start))
	ofile, err := os.Create("fractal_newton128.png")
	fmt.Println(time.Since(start))
	if err != nil {
		log.Fatal(err)
	}
	png.Encode(ofile, img)
	fmt.Println(time.Since(start))
}
