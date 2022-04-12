#include "stack.h"

int pop(Stack **s) {
  Stack *tail = *s;
  Stack *next = tail->next;
  *s = next;
  int ret = tail->value;
  delete tail;
  return ret;
}

void push(Stack **last, const int value) {
  Stack *target = new Stack{value, *last};
  *last = target;
}
