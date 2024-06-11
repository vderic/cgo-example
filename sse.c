#include <stdio.h>
#include <stdint.h>

#include <stdio.h>
#include <xmmintrin.h>
#include "sse.h"

int64_t sse(int n) {
    __m128 a = {1.0f, 2.0f, 3.0f, 4.0f};
    __m128 b = {1.1f, 2.1f, 3.1f, 4.1f};
    float c[4];
    __m128 ps = _mm_add_ps(a, b); // add
    _mm_storeu_ps(c, ps); // store

    printf("  source: (%5.1f, %5.1f, %5.1f, %5.1f)\n",
            a[0], a[1], a[2], a[3]);
    printf("  dest. : (%5.1f, %5.1f, %5.1f, %5.1f)\n",
           b[0], b[1], b[2], b[3]);
    printf("  result: (%5.1f, %5.1f, %5.1f, %5.1f)\n",
           c[0], c[1], c[2], c[3]);
    return 0;
}
