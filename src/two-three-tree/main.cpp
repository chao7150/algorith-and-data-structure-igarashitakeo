#include "two-three-tree.h"

int main() {
  Node *n = new Node{10, {-1, -1}, {nullptr, nullptr, nullptr}};
  Node *n2 = insert(n, 20);
  Node *n3 = insert(n2, 30);
  Node *n4 = insert(n3, 40);
  Node *n5 = insert(n4, 50);
  Node *n6 = insert(n5, 60);
}