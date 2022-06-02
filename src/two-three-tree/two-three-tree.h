#include <utility>

struct Node {
  int value = -1;
  int markers[2] = {-1, -1};
  Node *children[4];
};

int min(Node *node);
Node* insert(Node *node, const int value);
std::pair<Node *, Node *> _insert(Node *node, const int value);