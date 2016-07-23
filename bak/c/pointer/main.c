#include <stdio.h>

void
add(int *c, int a, int b)
{
    printf("a3=%x\n", &a);
    printf("b3=%x\n", &b);
    printf("c3=%x\n", c);
    *c = a + b;
}


int
main(int argc, char* argv[])
{
    int a = 0;
    int b = 0;
    int c = 0;
    a = 2;
    b = 3;
    printf("a1=%x\n", &a);
    printf("b1=%x\n", &b);
    printf("c1=%x\n", &c);
    add(&c, a, b);
    printf("%d\n", c);
    return 0;
}
