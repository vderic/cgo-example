#include <stdint.h>
#include "vectorize.h"

float gcc_dot_product(float *x, float *y, unsigned int n) {
    float res = 0;
    for (int i=0; i < n ; i++) {
	    res += x[i] * y[i];
    }
    return res;
}
