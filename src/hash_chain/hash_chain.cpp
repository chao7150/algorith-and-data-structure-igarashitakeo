#include "hash_chain.h"

int hash(const int value) { return value % 10; }

void insert(HashChain *h, const int value) {
  int index = hash(value);
  h->data[index].push_back(value);
}

bool member(HashChain *h, const int value) {
  int index = hash(value);
  for (int v : h->data[index]) {
    if (v == value) {
      return true;
    }
  }
  return false;
}
