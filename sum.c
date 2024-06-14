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

int fake_gcc_sum(unsigned int n) {
    float *xx = malloc(sizeof(float) * n);
    float *yy =  malloc(sizeof(float) * n);
    float *rr =  malloc(sizeof(float) * n);

    for (int i = 0 ; i < n ; i++) {
	    xx[i] = 2;
	    yy[i] = 3;
    }

    for (int i=0; i < n ; i++) {
	    rr[i] = xx[i] * yy[i];
    }

    //
    free(xx);
    free(yy);
    free(rr);
    return 100; 
}
