#include "wrapper.h"

/*
 *  Since the `PyObject_CallMethod` is a variadic function in the C
 *  library, create a wrapper around it for the number of arguments
 *  that our generator requires (0).
 */
PyObject*
random_generator(PyObject* module)
{
    PyObject* gen = PyObject_CallMethod(module, "random_generator", NULL);
    return gen;
}
