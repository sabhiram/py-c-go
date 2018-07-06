package main

/*
#cgo LDFLAGS: -lpython2.7
#cgo CFLAGS: -g -Wall

#include <python2.7/Python.h>
*/
import "C"

func main() {
	/*
	 *  Initialize and cleanup the python interpreter.
	 */
	C.Py_Initialize()
	defer C.Py_Finalize()
}
