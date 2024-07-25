#include "callbacks.h"
#include <pthread.h>
#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>

int callback_c(int data, void *opaque) { return callback_go(data, opaque); }

static int call_count = 0;

typedef struct {
  callback_type *cb;
  void *opaque;
} CallConf;

void *caller_thread(void *arg) {
  CallConf *callconf = arg;
  usleep(100 * 1);
  callconf->cb(++call_count, callconf->opaque);
  free(callconf);
  return 0;
}

int caller(callback_type cb, void *opaque) {
  pthread_t thread;

  CallConf *callconf = malloc(sizeof(CallConf));
  callconf->cb = cb;
  callconf->opaque = opaque;

  int result = pthread_create(&thread, NULL, &caller_thread, callconf);
  return result;
}
