#include <openssl/evp.h>

typedef void (*HashCallbackFn)(char* buf, int size);

void hashrequest(const char* input, int input_len, HashCallbackFn cb){
    cb((char*)input, input_len);
}