package main

/*
#cgo amd64 CFLAGS: -march=haswell -O3 -funroll-loops -ftree-vectorize -ffast-math
#cgo arm64 CFLAGS: -O3 -funroll-loops -ftree-vectorize
#include "avx.h"
#include "sse.h"
#include "vectorize.h"
#include "sse.hpp"
*/
import "C"
import (
	"fmt"
	"time"
	"github.com/vderic/cgo-example/peachpy"
)

func gosum(n int) int {
	result := 0
	for i := 0 ; i < n ; i++ {
		result += i
	}
	return result
}

func main() {
	start := time.Now()

	n := 5
	result := C.avx(C.int(n))

	diff := time.Since(start)

	fmt.Printf("Factorial of %d is %d\n", n, result)

	fmt.Printf("time = %d\n", diff)

	start = time.Now()
	result = C.sse(C.int(n))
	diff = time.Since(start)
	fmt.Printf("time = %d\n", diff)

	start = time.Now()
	result = C.sse2(C.int(n))
	diff = time.Since(start)
	fmt.Printf("time = %d\n", diff)

	n = 1000000000
	start = time.Now()
	result = C.gccsum(C.int(n))
	diff = time.Since(start)
	fmt.Printf("time = %d\n", diff)
	fmt.Printf("Vectorize Sum of %d is %d\n", n, result)

	n = 1000000000
	start = time.Now()
	sum := gosum(n)
	diff = time.Since(start)
	fmt.Printf("time = %d\n", diff)
	fmt.Printf("GoSum of %d = %d\n", n, sum)


	x := make([]float32, 2048)
        y := make([]float32, len(x))
        for i := 0; i < len(x); i++ {
                x[i] = 2.0
                y[i] = 3.0
        }
	z := peachpy.DotProduct(x, y, uint(len(x)))
}

