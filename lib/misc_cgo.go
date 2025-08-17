package lib

// #include <stdio.h>
// #include <stdlib.h>
// #include <unistd.h>
import "C"
import (
	"fmt"
	"unsafe"
)

func Putchar () {
	C.putchar('H')
	C.putchar('e')
	C.putchar('l')
	C.putchar('l')
	C.putchar('o')
	var c = int(C.getchar())
	fmt.Printf("c=%d\n", c)
}

func Unlink (s string) {
	cs := C.CString(s)
	defer C.free(unsafe.Pointer(cs))
	rc := int(C.unlink(cs))
	fmt.Printf("rc=%d\n", rc)
}