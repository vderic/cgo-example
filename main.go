package main

/*
#cgo amd64 CFLAGS: -march=haswell -O3 -funroll-loops -ftree-vectorize -ffast-math
#cgo arm64 CFLAGS: -O3 -funroll-loops -ftree-vectorize
#include "inc.h"
#include "avx.h"
#include "sse.h"
#include "vectorize.h"
#include "sse.hpp"
#include "sum.h"
*/
import "C"
import (
	"fmt"
	"time"
	"unsafe"
//	"github.com/vderic/cgo-example/peachpy"
)

func go_dot_product(x []float32, y []float32, n uint) float64 {
	result := float64(0)
	for i := 0 ; i < int(n) ; i++ {
		result += float64(x[i] * y[i])
	}
	return result
}

func go_sum(r []float32, x []float32, y []float32, n uint) {
	for i := 0 ; i < int(n) ; i++ {
		r[i] = x[i] * y[i]
	}
}

func go_sum_overhead(r []float32, x []float32, y []float32, n uint) {
	for i := 0 ; i < int(n) ; i++ {
		r[i] = multiply(x[i], y[i])
	}
}

//go:noinline
func multiply(x float32, y float32) float32 {
	return x * y
}

func inc(n int) int {
	for i := 0 ; i < 1000000000 ; i++ {
		n += 1
	}
	return n
}

func inc_overhead(n int) int {
	for i := 0 ; i < 1000000000 ; i++ {
		n = add(n)
	}
	return n
}

//go:noinline
func add(n int) int {
	return n+1
}


func main() {
	n := 100000000
	x := make([]float32, n)
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
	fmt.Printf("vectorized dot product: time = %d\n", diff)
	fmt.Printf("Vectorize dot product of is %f\n", gccresult)

	start = time.Now()
	goresult := go_dot_product(x, y, uint(len(x)))
	diff = time.Since(start)
	fmt.Printf("go dot proudct: time = %d\n", diff)
	fmt.Printf("Go dot product = %f\n", goresult)

	r := make([]float32, len(x))
	start = time.Now()
	C.gcc_sum((*C.float)(unsafe.Pointer(&r[0])), (*C.float)(unsafe.Pointer(&x[0])), (*C.float)(unsafe.Pointer(&y[0])),  C.uint(len(x)))
	diff1 := time.Since(start)
	fmt.Printf("vectorized sum: time = %d\n", diff1)

	start = time.Now()
	go_sum(r, x, y,  uint(len(x)))
	diff2 := time.Since(start)
	fmt.Printf("go sum: time = %d\n", diff2)

	start = time.Now()
	go_sum_overhead(r, x, y,  uint(len(x)))
	diff3 := time.Since(start)
	fmt.Printf("go sum overhead: time = %d\n", diff3)

	fmt.Printf("ratio compared to go: (GO/GO) 1 vs (GO_OVERHEAD/GO) %f vs (GCC/GO) %f\n",  float32(diff3)/float32(diff2), float32(diff1)/float32(diff2))

	start = time.Now()
	increment1 := inc(1000)
	diff1 = time.Since(start)
	fmt.Printf("go inc: time = %d, v=%d\n", diff1, increment1)

	start = time.Now()
	increment2 := inc_overhead(1000)
	diff2 = time.Since(start)
	fmt.Printf("inc overhead: time = %d, v=%d\n", diff2, increment2)

	start = time.Now()
	nn := 1000
	increment3 := C.gcc_inc(C.int(nn))
	diff3 = time.Since(start)
	fmt.Printf("gcc inc: time = %d, v=%d\n", diff3, increment3)

	fmt.Printf("ratio compared to go: (GO/GO) 1 vs (GO_OVERHEAD/GO) %f vs (GCC/GO) %f\n",  float32(diff2)/float32(diff1), float32(diff3)/float32(diff1))
	/*

	for i := 0 ; i < len(x) ; i++ {
		fmt.Printf("%f ", r[i])
	}
	fmt.Printf("\n")
	*/


	/* peachpy is buggy.  DO NOT USE IT */
	/*
	start = time.Now()
	result := peachpy.DotProduct(&x[0], &y[0], uint(len(x)))
	diff = time.Since(start)
	fmt.Printf("time = %d\n", diff)
	fmt.Printf("z = %f\n", result)
	*/


}



