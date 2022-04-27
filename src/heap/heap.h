struct Heap {
  int data[16];
  int last = 0;
};

void insert(Heap *h, const int value);
int pop(Heap *h);