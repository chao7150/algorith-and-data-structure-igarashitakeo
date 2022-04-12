struct Stack {
  int value;
  struct Stack *next;
};

int pop(Stack **last);
void push(Stack **last, const int value);
