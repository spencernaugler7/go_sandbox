package cgo

/*
#include <stdio.h>
#include <stdlib.h>

static int add(int a, int b) {
	return a + b;
}
*/
import "C"

import (
	"fmt"
	"io"
)

func Register(registry map[string]func(io.Writer)) map[string]func(io.Writer) {
	registry["cgo"] = Run
	return registry
}

func Run(writer io.Writer) {
	sum := C.add(C.int(3), C.int(4))
	fmt.Fprintln(writer, "sum (calculated in cgo)", int(sum))
}
