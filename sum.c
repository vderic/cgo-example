#include <stdint.h>
#include <stdio.h>
#include <stdlib.h>
#include "sum.h"

int gcc_sum(float *r, float *x, float *y, unsigned int n) {
    for (int i=0; i < n ; i++) {
	    r[i] = x[i] * y[i];
    }
    return 0;
}

/* When we modify the return float * r as return value, the processing time will increase greatly */
double fake_gcc_sum(float *r, float *x, float *y, unsigned int n) {
    double res = 0;

    for (int i=0; i < n ; i++) {
	    res += x[i] * y[i];
    }

    return res; 
}
