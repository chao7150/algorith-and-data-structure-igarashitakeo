#include <iostream>
#include <vector>

#include "binary_search.h"

int LENGTH = 4096;

int main() {
  std::vector<int> v;
  v.reserve(LENGTH);
  for (int i = 0; i < LENGTH; i++) {
    v.push_back(i);
  }
  for (int i = 0; i < LENGTH; i++) {
    auto [exists, loop] = binary_search(v, i);
    std::cout << i << "," << loop << std::endl;
  }
}