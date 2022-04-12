#include <gtest/gtest.h>

#include "stack.h"

TEST(stack, popSingle) {
  Stack *s = new Stack{1, nullptr};
  EXPECT_EQ(pop(&s), 1);
  Stack *t = new Stack{2, nullptr};
  EXPECT_EQ(pop(&t), 2);
}

TEST(stack, pushAndPop) {
  Stack *s = new Stack{1, nullptr};
  push(&s, 2);
  EXPECT_EQ(pop(&s), 2);
  EXPECT_EQ(pop(&s), 1);
}