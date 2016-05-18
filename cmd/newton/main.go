package main

import (
	"fmt"
	"image/png"
	"log"
	"os"
	"time"

	"github.com/ghurley/newton"
)

func main() {
	const upperLeft complex64 = -1.5 + 1.5i
	const lowerRight complex64 = 1.5 - 1.5i

	// Uncomment if you want profiling.
	/*
		f, err := os.Create("prog.prof")
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	*/

	start := time.Now()

	img := newton.DrawNewton64(640, 480, upperLeft, lowerRight)
	fmt.Println("Calculation time: ", time.Since(start))
	ofile, err := os.Create("fractal_newton.png")
	fmt.Println(time.Since(start))
	if err != nil {
		log.Fatal(err)
	}
	png.Encode(ofile, img)
	fmt.Println(time.Since(start))
}
