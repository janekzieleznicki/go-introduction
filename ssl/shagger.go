package ssl

import (
  "unsafe"
  "encoding/hex"
)

/*
#cgo LDFLAGS: -lssl -lcrypto
#include <openssl/evp.h>
*/
import "C"

func Sha256(input string) string {
  var context *C.EVP_MD_CTX = C.EVP_MD_CTX_new()

  C.EVP_DigestInit_ex(context, C.EVP_sha256(), nil)

  //This copies the content to CString
  data := C.CString(input)
  defer C.free(unsafe.Pointer(data))

  C.EVP_DigestUpdate(context, (unsafe.Pointer(data)), (C.ulong) (len(input)))

  md := [C.EVP_MAX_MD_SIZE]byte{}
  var md_len C.uint
  
  C.EVP_DigestFinal_ex(context, (*C.uchar) (unsafe.Pointer(&md)), &md_len)

  //This copies the content from md to return string
  return hex.EncodeToString(md[:md_len])
}