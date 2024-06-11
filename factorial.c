#include <stdio.h>
#include <immintrin.h>
#include "factorial.h"

int64_t hello(int n) {
    __m256 a = {1.0f, 2.0f, 3.0f, 4.0f, 5.0f, 6.0f, 7.0f, 8.0f};
    __m256 b = {1.1f, 2.1f, 3.1f, 4.1f, 5.1f, 6.1f, 7.1f, 8.1f};
    __m256 c;

    c = _mm256_add_ps(a, b);

    printf("  source: (%5.1f, %5.1f, %5.1f, %5.1f, %5.1f, %5.1f, %5.1f, %5.1f)\n",
            a[0], a[1], a[2], a[3], a[4], a[5], a[6], a[7]);
    printf("  dest. : (%5.1f, %5.1f, %5.1f, %5.1f, %5.1f, %5.1f, %5.1f, %5.1f)\n",
            b[0], b[1], b[2], b[3], b[4], b[5], b[6], b[7]);
    printf("  result: (%5.1f, %5.1f, %5.1f, %5.1f, %5.1f, %5.1f, %5.1f, %5.1f)\n",
           c[0], c[1], c[2], c[3], c[4], c[5], c[6], c[7]);
    return 0;
}

/*
#include <stdio.h>
#include <stdint.h>

#include <stdio.h>
#include <xmmintrin.h>

int64_t hello(int n) {
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
*/

/*
int64_t hello(int n) {
    if (n == 0) {
        return 1;
    } else {
        return n * hello(n - 1);
    }
}
*/
