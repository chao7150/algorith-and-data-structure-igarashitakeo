#include "binary_search_tree.h"

void insert(Node *h, const int value) {
  if (value <= h->value) {
    if (h->left == nullptr) {
      h->left = new Node{value, nullptr, nullptr};
    } else {
      insert(h->left, value);
    }
  } else {
    if (h->right == nullptr) {
      h->right = new Node{value, nullptr, nullptr};
    } else {
      insert(h->right, value);
    }
  }
};

bool search(Node *h, const int value) {
  if (h == nullptr) {
    return false;
  }
  if (h->value == value) {
    return true;
  }
  if (value < h->value) {
    return search(h->left, value);
  }
  return search(h->right, value);
}

Node *remove(Node *h, const int value) {
  if (h->value == value) {
    if (h->left == nullptr && h->right == nullptr) {
      delete h;
      return nullptr;
    } else if (h->right == nullptr) {
      h->left = remove(h->left, value);
    } else if (h->left == nullptr) {
      h->right = remove(h->right, value);
    } else {
      h->value = findMin(h->right);
      h->right = remove(h->right, h->value);
    }
  }
  else {
    if (value < h->value) {
      h->left = remove(h->left, value);
    } else {
      h->right = remove(h->right, value);
    }
  }
  return h;
}

int findMin(Node *h) {
  if (h->left == nullptr) {
    return h->value;
  }
  return findMin(h->left);
}