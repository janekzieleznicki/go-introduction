package ssl

import (
  "fmt"
)

/*
#cgo LDFLAGS: -lssl -lcrypto -L -lhasher
#include <openssl/evp.h>

typedef void (*HashCallbackFn)(char* buf, int size);

extern void goHashCallback(char*, int);

// this is implemented in hasher.c
void hashrequest(const char* input, int input_len, HashCallbackFn);
*/
import "C"

//export goHashCallback
func goHashCallback(data *C.char, data_len C.int){
  s := C.GoString(data)
  fmt.Printf("calledback with %s", s)
}

func Sha256Calledback(input string) {
  //This copies the content to CString
  data := C.CString(input)

  C.hashrequest( data, (C.int) (len(input)), C.HashCallbackFn(C.goHashCallback) )
}

