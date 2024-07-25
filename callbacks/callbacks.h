
typedef int(callback_type)(int, void *);

extern int callback_go(int data, void *opaque);

int callback_c(int data, void *opaque);

int caller(callback_type cb, void *opaque);
