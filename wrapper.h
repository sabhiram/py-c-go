#ifndef _WRAPPER_H_
#define _WRAPPER_H_

#include <python2.7/Python.h>

PyObject* New(void);
const int Next(PyObject* generator);

#endif // _WRAPPER_H_
