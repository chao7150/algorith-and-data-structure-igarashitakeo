#include "list.h"

void push(Cell *cell, const int value) {
  if (cell->next == nullptr) {
    Cell *tail = new Cell{value, nullptr};
    cell->next = tail;
  } else {
    push(cell->next, value);
  }
};

int at(Cell *cell, const int index) {
  if (index == 0) {
    return cell->value;
  }
  return at(cell->next, index - 1);
};

int length(const Cell cell) {
  int itr = 1;
  Cell pos = cell;
  while (pos.next != nullptr) {
    pos = *pos.next;
    itr++;
  }
  return itr;
}

void pop_back(Cell *cell) {
  if (cell->next == nullptr) {
    return;
  }
  if (cell->next->next == nullptr) {
    delete cell->next->next;
    cell->next = nullptr;
    return;
  }
  pop_back(cell->next);
}

Cell *at_pointer(Cell *cell, const int position) {
  Cell *current = cell;
  for (int i = 0; i < position; i++) {
    cell = cell->next;
  }
  return cell;
}

void insert(Cell **head, const int position, const int value) {
  Cell *target = new Cell{value, nullptr};
  if (position == 0) {
    target->next = *head;
    *head = target;
  }
  Cell *previous = at_pointer(*head, position - 1);
  Cell *next = previous->next;
  previous->next = target;
  target->next = next;
};

void erase(Cell **head, const int position) {
  Cell *phead = *head;
  if (position == 0) {
    // 削除対象は先頭である
    *head = phead->next;
    delete phead;
    return;
  }
  if (position == 1) {
    Cell *ptarget = phead->next;
    Cell *pchild = ptarget->next;
    if (pchild == nullptr) {
      // 削除対象は末尾要素である
      delete phead->next;
      phead->next = nullptr;
    } else {
      // 削除対象の後ろにも要素がある
      delete phead->next;
      phead->next = pchild;
    }
    return;
  }
  erase(&(phead->next), position - 1);
}