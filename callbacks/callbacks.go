package main

/*
#include "callbacks.h"
*/
import "C"
import (
	"fmt"
	"runtime"
	"time"
	"unsafe"
)

type ComplexData struct {
	intData    int
	stringData string
}

//export callback_go
func callback_go(data C.int, opaque unsafe.Pointer) C.int {
	complexData := (*ComplexData)(opaque)
	fmt.Printf("Callback called %d, with data:\t%v\n", data, complexData)
	return 0
}
//export struct_callback_go 
func struct_callback_go(val *C.struct_AsStruct) {

}

//export typedef_callback_go
func typedef_callback_go(val *C.Typedefed){

}

func main() {
	complexData := new(ComplexData)
	complexData.intData = 123
	complexData.stringData = "Here is some massive data in complex Data allocated on heap"

	stackData := ComplexData{
		intData:    321,
		stringData: "Here is some other data allocated somewhere else",
	}
	pnr := new(runtime.Pinner)
	pnr.Pin(complexData)
	pnr.Pin(&stackData)
	C.caller((*C.callback_type)(C.callback_c), unsafe.Pointer(complexData))
	C.caller((*C.callback_type)(C.callback_c), unsafe.Pointer(&stackData))
	// pnr = nil
	time.Sleep(100 * time.Microsecond)
	// pnr.Unpin()
}
