struct Node {
  int value;
  Node *left;
  Node *right;
};

void insert(Node *h, const int value);
bool search(Node *h, const int value);
Node *remove(Node *h, const int value);
int findMin(Node *h);
