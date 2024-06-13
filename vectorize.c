#include <stdint.h>
#include "vectorize.h"

double gcc_dot_product(float *x, float *y, unsigned int n) {
    double res = 0;
    for (int i=0; i < n ; i++) {
	    res += x[i] * y[i];
    }
    return res;
}
