#define BUCKET_SIZE 10

struct HashOpenAddress {
  int data[BUCKET_SIZE] = {-1, -1, -1, -1, -1, -1, -1, -1, -1, -1};
};

bool member(HashOpenAddress *h, const int value);
void insert(HashOpenAddress *h, const int value);
