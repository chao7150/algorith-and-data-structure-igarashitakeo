#include "heap.h"

int main() {
  Heap *h = new Heap{{0, 1, 2, 3}, 3};
  pop(h);
}