#include "queue.h"

int at(Queue q, const int position) {
  return q.data[(q.top + position) % q.length];
};

void push(Queue *q, const int value) {
  q->data[q->bottom % q->length] = value;
  q->bottom++;
}

int pop(Queue *q) {
  int ret = at(*q, 0);
  q->top++;
  return ret;
};
