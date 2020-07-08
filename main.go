package main

/*
#include <stdio.h>
#include <stdlib.h>

static void printHello() {
	printf("hello");
}
*/
import "C"

import (
	"fmt"
)

func main() {
	// // Glob - Gets the plugin to be loaded
	// plugins, err := filepath.Glob("print.so")
	// if err != nil {
	// 	panic(err)
	// }
	// Open - Loads the plugin
	fmt.Println("Loading plugin print.so")
	// module = "print.so"

	// mod := C.CString(module)
	// defer C.free(unsafe.Pointer(mod))

	// p := C.New(mod)

	C.printHello()
}
