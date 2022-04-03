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
