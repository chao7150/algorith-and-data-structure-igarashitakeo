#include "hash_open_address.h"

int hash(const int value) { return value % BUCKET_SIZE; }

bool member(HashOpenAddress *h, const int value) {
  int index = hash(value);
  for (int i = 0; i < BUCKET_SIZE; i++) {
    if (h->data[(index + i) % BUCKET_SIZE] == -1)
      return false;
    else 
  }
  return false;
}

bool _insert(HashOpenAddress *h, const int value, const int index,
             const bool insert = false) {
  bool free = h->data[index] == -1;
  if (insert)
    h->data[index] = value;
  return free;
}

void insert(HashOpenAddress *h, const int value) {
  int index = hash(value);
  for (int i = 0; i < BUCKET_SIZE; i++) {
    if (_insert(h, value, (index + i) % BUCKET_SIZE)) {
      _insert(h, value, (index + i) % BUCKET_SIZE, true);
      return;
    }
  }
}
