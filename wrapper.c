#include <stdio.h>
#include <stdlib.h>
#include <assert.h>
#include <string.h>

#include "wrapper.h"

/*
 *  New loads the `generator` module and returns a PyObject*
 *  which represents the `random_generator` generator.
 */
PyObject*
New(void)
{
    /*
     *  Import the python module from the file `generator.py`.
     */
    PyObject* module = PyImport_ImportModule("generator");
    if (module == NULL)
    {
        printf("generator import not found\n");
        return NULL;
    }

    PyObject* gen = PyObject_CallMethod(module, "random_generator", NULL);

    /*
     *  Remove the reference to the module.
     *  TODO: Verify if this needs to be deferred?
     */
    Py_XDECREF(module);

    return gen;
}

/*
 *  Returns the next random value between [0, 100].
 *
 *  Returns -1 if empty.
 */
const int
Next(PyObject* gen)
{
    PyObject* line = PyIter_Next(gen);
    if (line == NULL)
        return -1;

    int v = PyInt_AsLong(line);

    /*
     *  The line has been used, relinquish!
     */
    Py_DECREF(line);

    return v;
}
