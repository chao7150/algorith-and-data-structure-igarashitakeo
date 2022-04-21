#include <gtest/gtest.h>

#include "queue.h"

TEST(queue, access) {
  struct Queue q = {{0}};
  EXPECT_EQ(at(q, 0), 0);
}

TEST(queue, push) {
  Queue *q = new Queue;
  push(q, 0);
  EXPECT_EQ(at(*q, 0), 0);
  push(q, 1);
  EXPECT_EQ(at(*q, 1), 1);
  push(q, 2);
  EXPECT_EQ(at(*q, 2), 2);
}

TEST(queue, pop) {
  Queue *q = new Queue;
  push(q, 0);
  push(q, 1);
  EXPECT_EQ(pop(q), 0);
  EXPECT_EQ(at(*q, 0), 1);
}

TEST(queue, overflow) {
  Queue *q = new Queue;
  for (int i = 0; i < 16; i++) {
    push(q, i);
  }
  EXPECT_EQ(at(*q, 0), 0);
  EXPECT_EQ(at(*q, 15), 15);
  EXPECT_EQ(pop(q), 0);
  push(q, 16);
  EXPECT_EQ(at(*q, 15), 16);
  EXPECT_EQ(q->data[0], 16);
}