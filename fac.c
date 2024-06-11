#include <stdint.h>
#include "fac.h"

int64_t fac(int n) {
    int64_t sum = 0;
    #pragma clang loop vectorize(enable)
    for (int i=0; i < n ; i++) {
	    sum += i;
    }
    return sum;
}
