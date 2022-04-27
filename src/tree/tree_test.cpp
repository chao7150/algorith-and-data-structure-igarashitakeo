#include <gtest/gtest.h>

#include "tree.h"

TEST(tree, hasValue) {
  Tree *t = new Tree{0, nullptr, nullptr};
  EXPECT_EQ(t->value, 0);
}
