package main

/*
#cgo LDFLAGS:
#cgo CFLAGS: -g -Wall

#include <stdio.h>

void print_x(int x)
{
    printf("X = %d\n", x);
}
*/
import "C"

func main() {
	C.print_x(1)
	C.print_x(2)
	C.print_x(3)
}
