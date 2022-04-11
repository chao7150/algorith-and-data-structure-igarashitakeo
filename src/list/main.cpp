#include "list.h"
#include <iostream>

void _p(Cell *h) {
  Cell tail = {2, nullptr};
  h->next = &tail;
}

void p(Cell *h) {
  Cell *tail = new Cell;
  tail->value = 2;
  tail->next = nullptr;
  h->next = tail;
}

void scenarioA() {
  Cell c1 = {0, nullptr};
  p(&c1);
  std::cout << at(&c1, 0) << std::endl;
  std::cout << c1.next->value << std::endl;
  std::cout << at(&c1, 1) << std::endl;
}

void scenarioB() {
  Cell c1 = {0, nullptr};
  Cell tail = {2, nullptr};
  c1.next = &tail;
  std::cout << at(&c1, 0) << std::endl;
  std::cout << c1.next->value << std::endl;
  std::cout << at(&c1, 1) << std::endl;
}

void scenarioC() {
  Cell c = {1, nullptr};
  Cell *ptr = &c;
  push(&c, 2);
  insert(&ptr, 0, 0);
}

void scenarioD() {
  Cell *ptr = new Cell{1, nullptr};
  push(ptr, 2); // [1, 2]
  push(ptr, 3); // [1, 2, 3]
  erase(&ptr, 0);
  erase(&ptr, 0);
}

int main() {
  // scenarioA();
  // scenarioB();
  // scenarioC();
  scenarioD();
}
