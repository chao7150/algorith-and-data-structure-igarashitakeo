#include <gtest/gtest.h>

#include "list.h"

TEST(cell, assign) {
  Cell c1 = {1, nullptr};
  EXPECT_EQ(c1.value, 1);
  c1.value = 2;
  EXPECT_EQ(c1.value, 2);
}

TEST(cell, accessandpush) {
  Cell c2 = {1, nullptr};
  EXPECT_EQ(at(&c2, 0), 1);
  push(&c2, 2);
  EXPECT_EQ(at(&c2, 1), 2);
  push(&c2, 3);
  EXPECT_EQ(at(&c2, 2), 3);
}

TEST(cell, length) {
  Cell c = {0, nullptr};
  EXPECT_EQ(length(c), 1);
  push(&c, 1);
  EXPECT_EQ(length(c), 2);
  push(&c, 2);
  EXPECT_EQ(length(c), 3);
}

TEST(cell, pop_back) {
  Cell c = {0, nullptr};
  push(&c, 1);
  pop_back(&c);
  EXPECT_EQ(length(c), 1);
}

TEST(cell, insertHead) {
  Cell c = {1, nullptr};
  Cell *ptr = &c;
  push(&c, 2);
  insert(&ptr, 0, 0);
  EXPECT_EQ(at(ptr, 0), 0);
}

TEST(cell, insert) {
  Cell c = {1, nullptr};
  Cell *ptr = &c;
  push(&c, 2);        // [1, 2]
  insert(&ptr, 1, 0); // [1, 0, 2]
  EXPECT_EQ(at(ptr, 1), 0);
  insert(&ptr, 1, 3); // [1, 3, 0, 2]
  EXPECT_EQ(at(ptr, 1), 3);
  insert(&ptr, 4, 4);
  EXPECT_EQ(at(ptr, 4), 4);
}

TEST(cell, eraseHead) {
  Cell *ptr = new Cell{1, nullptr};
  push(ptr, 2); // [1, 2]
  push(ptr, 3); // [1, 2, 3]
  erase(&ptr, 0);
  EXPECT_EQ(at(ptr, 0), 2);
  erase(&ptr, 0);
  EXPECT_EQ(at(ptr, 0), 3);
}

TEST(cell, erase) {
  Cell *ptr = new Cell{1, nullptr};
  push(ptr, 2);   // [1, 2]
  push(ptr, 3);   // [1, 2, 3]
  erase(&ptr, 1); // [1, 3]
  EXPECT_EQ(at(ptr, 1), 3);
  erase(&ptr, 1); // [1]
  EXPECT_EQ(length(*ptr), 1);
}