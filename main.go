package main

// #cgo LDFLAGS: -lpython2.7
// #cgo CFLAGS: -g -Wall
// #include "wrapper.h"
import "C"
import (
	"fmt"
	"os"
)

func main() {
	/*
	 *  Initialize and cleanup the python interpreter.
	 */
	C.Py_Initialize()
	defer C.Py_Finalize()

	/*
	 *  Create a new generator from the C codez.
	 */
	gen := C.NewGenerator()
	if gen == nil {
		fmt.Printf("Fatal error: generator is null!\n")
		os.Exit(1)
	}

	for n := C.Next(gen); n >= 0; n = C.Next(gen) {
		fmt.Printf("Next random number: %d\n", n)
	}
}
