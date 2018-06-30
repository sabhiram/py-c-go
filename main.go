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
	 *  Create a new generator using the C wrapper.
	 */
	genObj := C.New()
	if genObj == nil {
		fmt.Printf("Fatal error: generator is null!\n")
		os.Exit(1)
	}

	/*
	 *  Generate numbers!
	 */
	for n := C.Next(genObj); n >= 0; n = C.Next(genObj) {
		fmt.Printf("Next random number: %d\n", n)
	}
}
