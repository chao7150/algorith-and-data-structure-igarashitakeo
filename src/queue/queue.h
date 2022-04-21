struct Queue {
  int data[16];
  int top = 0;
  int bottom = 0;
  int length = 16;
};

int at(Queue q, const int position);
void push(Queue *q, const int value);
int pop(Queue *q);
