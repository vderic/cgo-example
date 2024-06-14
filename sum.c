#include <stdint.h>
#include "sum.h"

int gcc_sum(float *r, float *x, float *y, unsigned int n) {
    for (int i=0; i < n ; i++) {
	    r[i] = x[i] * y[i];
    }
    return 0;
}
