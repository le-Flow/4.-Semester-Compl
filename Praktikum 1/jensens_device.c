#include <stdio.h>

#define sum(i, max, expr, r) {\
    for((i) = 0; (i) < (max); (i)++){\
        (r) += (expr);\
    }\
}
    
int main() {
    puts("Jensens Device");
    
    int i = 0;
    int x = 10;
    int r = 0;
    
    sum(i, 3, i * x, r);
    
    printf("sum(i * x) = %d\n", r);
    return 0;
}
