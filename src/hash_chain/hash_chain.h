#include <vector>

struct HashChain {
  std::vector<int> data[10];
};

void insert(HashChain *h, const int value);
bool member(HashChain *h, const int value);
void delete_element(HashChain *h, const int value);
