#include "heap.h"
#include <iostream>

void insert(Heap *h, const int value) {
  h->last++;
  h->data[h->last] = value;
  int p = h->last;
  while (p != 0) {
    int parent_index = (p - 1) / 2;
    int current = h->data[p];
    int parent = h->data[parent_index];
    if (current < parent) {
      h->data[p] = parent;
      h->data[parent_index] = current;
    }
    p = parent_index;
  };
};

void swap_by_index(int l[], const int x, const int y) {
  int tmp = l[x];
  l[x] = l[y];
  l[y] = tmp;
}

int pop(Heap *h) {
  int ret = h->data[0];
  h->data[0] = h->data[h->last];
  h->last--;
  int p = 0;
  while (p <= (h->last - 1) / 2) {
    int self = h->data[p];
    int left_child = h->data[p * 2 + 1];
    int right_child = h->data[p * 2 + 2];
    if (self < left_child && self < right_child) {
      break;
    }
    int smaller_child_index = left_child < right_child ? p * 2 + 1 : p * 2 + 2;
    swap_by_index(h->data, p, smaller_child_index);
    p = smaller_child_index;
  }
  return ret;
}