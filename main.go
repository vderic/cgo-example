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
	"unsafe"
	"github.com/vderic/cgo-example/peachpy"
)

func go_dot_product(x []float32, y []float32, n uint) float64 {
	result := float64(0)
	for i := 0 ; i < int(n) ; i++ {
		result += float64(x[i] * y[i])
	}
	return result
}

func main() {
	x := make([]float32, 204800000)
        y := make([]float32, len(x))
        for i := 0; i < len(x); i++ {
                x[i] = 2.0
                y[i] = 3.0
        }

	start := time.Now()
	diff := time.Since(start)

	/*
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
	*/

	start = time.Now()
	gccresult := C.gcc_dot_product((*C.float)(unsafe.Pointer(&x[0])), (*C.float)(unsafe.Pointer(&y[0])), C.uint(len(x)))
	diff = time.Since(start)
	fmt.Printf("time = %d\n", diff)
	fmt.Printf("Vectorize Sum of is %f\n", gccresult)

	start = time.Now()
	goresult := go_dot_product(x, y, uint(len(x)))
	diff = time.Since(start)
	fmt.Printf("time = %d\n", diff)
	fmt.Printf("GoSum = %f\n", goresult)


	start = time.Now()
	result := peachpy.DotProduct(&x[0], &y[0], uint(len(x)))
	diff = time.Since(start)
	fmt.Printf("time = %d\n", diff)
	fmt.Printf("z = %f\n", result)
}

