#include <gtest/gtest.h>
#include <iostream>

#include "heap.h"

TEST(heap, assign) {
  Heap *h = new Heap{{0, 1, 2, 3}, 3};
  insert(h, 4);
  EXPECT_EQ(h->data[4], 4);
  insert(h, 1);
  EXPECT_EQ(h->data[2], 1);
  EXPECT_EQ(h->data[h->last], 2);
}

TEST(heap, pop) {
  Heap *h = new Heap{{0, 1, 2, 3}, 3};
  EXPECT_EQ(pop(h), 0);
  EXPECT_EQ(h->data[0], 1);
  EXPECT_EQ(h->data[1], 3);
  EXPECT_EQ(h->data[2], 2);
}