
typedef int(callback_type)(int, void *);

extern int callback_go(int data, void *opaque);

int callback_c(int data, void *opaque);

int caller(callback_type cb, void *opaque);


struct AsStruct {
  int a;
  int b;
  int c;
};
typedef struct {
  int a;
  int b;
  int c;
} Typedefed;


extern void struct_callback_go(struct AsStruct * val);
extern void typedef_callback_go(Typedefed * val);