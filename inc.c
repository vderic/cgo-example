#include <stdint.h>
#include "inc.h"

int gcc_inc(int n) {
    for (int i=0; i < 1000000000 ; i++) {
	    n += 1;
    }
    return n;
}
