#ifndef _WRAPPER_H_
#define _WRAPPER_H_

#include <python2.7/Python.h>

typedef struct
{
    PyObject* instance;
} Generator_t;

Generator_t* NewGenerator(void);
const int Next(const Generator_t* generator);

#endif // _WRAPPER_H_
