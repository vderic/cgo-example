package main

/*
#cgo amd64 CFLAGS: -march=haswell -O3 -funroll-loops -ftree-vectorize -ffast-math
#cgo arm64 CFLAGS: -O3 -funroll-loops -ftree-vectorize
#include "factorial.h"
#include "sse.h"
#include "fac.h"
#include "sse.hpp"
*/
import "C"
import (
	"fmt"
	"time"
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
	result := C.hello(C.int(n))

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
	result = C.fac(C.int(n))
	diff = time.Since(start)
	fmt.Printf("time = %d\n", diff)
	fmt.Printf("Vectorize Sum of %d is %d\n", n, result)

	n = 1000000000
	start = time.Now()
	sum := gosum(n)
	diff = time.Since(start)
	fmt.Printf("time = %d\n", diff)
	fmt.Printf("GoSum of %d = %d\n", n, sum)

}

