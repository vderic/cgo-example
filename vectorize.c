#include <stdint.h>
#include "vectorize.h"

int64_t gccsum(int n) {
    int64_t sum = 0;
    #pragma clang loop vectorize(enable)
    for (int i=0; i < n ; i++) {
	    sum += i;
    }
    return sum;
}
