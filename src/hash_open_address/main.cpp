#include "hash_open_address.h"

int main() {
  HashOpenAddress *h = new HashOpenAddress;
  insert(h, 5);
  insert(h, 15);
  insert(h, 25);
}
