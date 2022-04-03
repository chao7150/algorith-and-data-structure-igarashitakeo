struct Cell {
  int value;
  struct Cell* next;
};

void push(Cell *cell, const int value);
int at(Cell *cell, const int index);
int length(const Cell cell);
void pop_back(Cell* cell);
