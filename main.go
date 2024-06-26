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
	"sync"
	"time"
	"unsafe"
	// "github.com/vderic/cgo-example/peachpy"
	"github.com/gonum/floats"
)

func go_dot_product(x []float32, y []float32, n uint) float64 {
	result := float64(0)
	for i := 0; i < int(n); i++ {
		result += float64(x[i] * y[i])
	}
	return result
}

func go_sum(r []float32, x []float32, y []float32, n uint) {
	for i := 0; i < int(n); i++ {
		r[i] = x[i] * y[i]
	}
}

func go_sum_overhead(r []float32, x []float32, y []float32, n uint, fn func(x, y float32) float32) {
	for i := 0; i < int(n); i++ {
		r[i] = fn(x[i], y[i])
	}
}

func multiply(x float32, y float32) float32 {
	return x * y
}

func gor_sum(r []float32, x []float32, y []float32, n uint) {
	var wg sync.WaitGroup
	nthread := 2
	chunkSize := int(n) / nthread
	for i := 0; i < nthread; i++ {
		wg.Add(1)
		go func(start, end int) {
			defer wg.Done()
			for j := start; j < end; j++ {
				r[j] = x[j] * y[j]
			}
		}(i*chunkSize, (i+1)*chunkSize)
	}
	wg.Wait()
}

func inc(n int) int {
	for i := 0; i < 1000000000; i++ {
		n += 1
	}
	return n
}

func inc_overhead(n int, fn func(int, int) int) int {
	j := 1
	for i := 0; i < 1000000000; i++ {
		n = fn(n, j)
	}
	return n
}

func add(n int, m int) int {
	return n + m
}

func main() {
	n := 100000000
	x := make([]float32, n)
	y := make([]float32, len(x))
	r := make([]float32, len(x))
	for i := 0; i < len(x); i++ {
		x[i] = 2.0
		y[i] = 3.0
	}

	start := time.Now()

	/*
		diff := time.Since(start)
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
	*/

	start = time.Now()
	gccresult := C.gcc_dot_product((*C.float)(unsafe.Pointer(&x[0])), (*C.float)(unsafe.Pointer(&y[0])), C.uint(len(x)))
	diff7 := time.Since(start)
	//fmt.Printf("vectorized dot product: time = %d\n", diff7)
	//fmt.Printf("Vectorize dot product of is %f\n", gccresult)

	start = time.Now()
	goresult := go_dot_product(x, y, uint(len(x)))
	diff8 := time.Since(start)

	fmt.Printf("dot product result = %f %f\n", gccresult, goresult)
	//fmt.Printf("go dot proudct: time = %d\n", diff8)
	//fmt.Printf("Go dot product = %f\n", goresult)

	fmt.Printf("Dot Product function\n")
	fmt.Printf("       |    GO            |     GCC       |\n")
	fmt.Printf(" Time  |    %d     |     %d        |\n", diff8, diff7)
	fmt.Printf("ratio compared to go: (GO/GO) 1 vs (GCC/GO) %f\n", float32(diff7)/float32(diff8))

	start = time.Now()
	fakev := C.fake_gcc_sum((*C.float)(unsafe.Pointer(&r[0])), (*C.float)(unsafe.Pointer(&x[0])), (*C.float)(unsafe.Pointer(&y[0])), C.uint(len(x)))
	diff6 := time.Since(start)
	fmt.Printf("single return value v = %f\n", float64(fakev))
	//fmt.Printf("fake gcc sum %d\n", diff6)

	start = time.Now()
	gor_sum(r, x, y, uint(len(x)))
	diff4 := time.Since(start)
	//fmt.Printf("go routine sum: time = %d\n", diff4)

	x64 := make([]float64, n)
	y64 := make([]float64, len(x))
	r64 := make([]float64, len(x))
	start = time.Now()
	floats.AddTo(r64, x64, y64)
	diff5 := time.Since(start)

	/*
		for i := 0 ; i < 10 ; i++ {
			fmt.Printf(" %f ", r[i])
		}
		fmt.Print()
	*/
	fmt.Printf("gonum %d\n", diff5)

	start = time.Now()
	C.gcc_sum((*C.float)(unsafe.Pointer(&r[0])), (*C.float)(unsafe.Pointer(&x[0])), (*C.float)(unsafe.Pointer(&y[0])), C.uint(len(x)))
	diff1 := time.Since(start)
	//fmt.Printf("vectorized sum: time = %d\n", diff1)

	start = time.Now()
	go_sum(r, x, y, uint(len(x)))
	diff2 := time.Since(start)
	//fmt.Printf("go sum: time = %d\n", diff2)

	start = time.Now()
	go_sum_overhead(r, x, y, uint(len(x)), func(x, y float32) float32 { return x * y })
	diff3 := time.Since(start)
	//fmt.Printf("go sum overhead: time = %d\n", diff3)

	//fmt.Printf("r[0] = %f\n", r[0])

	fmt.Printf("Mutliply function\n")
	fmt.Printf("       |    GO            |  GO OVERHEAD          |    GCC       |  Go Routine   | GCC Single return value\n")
	fmt.Printf(" Time  |    %d     |      %d       |    %d        |      %d      |      %d     \n", diff2, diff3, diff1, diff4, diff6)
	fmt.Printf("ratio compared to go: (GO/GO) 1 vs (GO_OVERHEAD/GO) %f vs (GCC/GO) %f vs (Go Rountine/Go) %f vs (FAKEGCC/GO) %f\n",
		float32(diff3)/float32(diff2), float32(diff1)/float32(diff2), float32(diff4)/float32(diff2), float32(diff6)/float32(diff2))

	start = time.Now()
	increment1 := inc(1000)
	diff1 = time.Since(start)
	//fmt.Printf("go inc: time = %d, v=%d\n", diff1, increment1)

	start = time.Now()
	increment2 := inc_overhead(1000, func(n int, m int) int { return n + m })
	diff2 = time.Since(start)
	//fmt.Printf("inc overhead: time = %d, v=%d\n", diff2, increment2)

	start = time.Now()
	increment3 := int(C.gcc_inc(C.int(1000)))
	diff3 = time.Since(start)
	//fmt.Printf("gcc inc: time = %d, v=%d\n", diff3, increment3)

	fmt.Printf("%d %d %d\n", increment1, increment2, increment3)

	fmt.Printf(" INC function\n")
	fmt.Printf("       |    GO            |  GO OVERHEAD          |    GCC     \n")
	fmt.Printf(" Time  |    %d     |      %d       |    %d\n", diff1, diff2, diff3)
	//fmt.Printf("  Ratio |   %f      |  %f           |    %f\n", 1.0, float32(diff2)/float32(diff1), float32(diff3)/float32(diff1))

	fmt.Printf("ratio compared to go: (GO/GO) 1 vs (GO_OVERHEAD/GO) %f vs (GCC/GO) %f\n", float32(diff2)/float32(diff1), float32(diff3)/float32(diff1))
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
