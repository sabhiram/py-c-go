#include <stdio.h>
#include <stdlib.h>
#include <assert.h>
#include <string.h>

#include "wrapper.h"

/*
 *  This method creates an returns a `Generator_t*` which wraps the
 *  underlying python `fuzz.Fuzz()` instance.
 */
Generator_t*
NewGenerator(void)
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

    PyObject* rgen = PyObject_CallMethod(module, "random_generator", NULL);
    if (rgen == NULL)
    {
        printf("random_generator not found\n");
        return NULL;
    }

    Generator_t* gen = malloc(sizeof *gen);
    gen->instance = rgen;
    return gen;
}

/*
 *  Returns the next random value between [0, 100].
 *
 *  Returns -1 if empty.
 */
const int
Next(const Generator_t* gen)
{
    PyObject* line = PyIter_Next(gen->instance);
    if (line == NULL)
        return -1;

    int v = PyInt_AsLong(line);
    Py_DECREF(line);

    return v;
}
