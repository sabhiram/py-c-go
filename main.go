package main

// #cgo LDFLAGS: -lpython2.7
// #cgo CFLAGS: -g -Wall
// #include "wrapper.h"
import "C"
import (
	"fmt"
	"os"
	"unsafe"
)

func main() {
	/*
	 *  Initialize and cleanup the python interpreter.
	 */
	C.Py_Initialize()
	defer C.Py_Finalize()

	/*
	 *  Import the python module from the file `generator.py`.
	 */
	module := C.PyImport_ImportModule(C.CString("generator"))
	if unsafe.Pointer(module) == nil {
		fmt.Printf("Unable to import `generator.py`\n")
		os.Exit(1)
	}

	/*
	 *  Create a new generator using the C wrapper.
	 */
	gen := C.random_generator(module)
	if gen == nil {
		fmt.Printf("Fatal error: generator is null!\n")
		os.Exit(1)
	}

	/*
	 *  Generate numbers!
	 */
	for l := C.PyIter_Next(gen); unsafe.Pointer(l) != nil; l = C.PyIter_Next(gen) {
		fmt.Printf("Next random number: %d\n", C.PyInt_AsLong(l))
	}
}
