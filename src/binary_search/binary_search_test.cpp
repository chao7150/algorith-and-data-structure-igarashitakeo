#include <gtest/gtest.h>
#include <vector>

#include "binary_search.h"

std::vector<int> v = {1, 3, 5};

TEST(slice_vector, a) {
  std::vector<int> v1 = {1, 3, 5, 7};
  EXPECT_TRUE(v == slice_vector(v1, 0, 2));
  std::vector<int> v2 = {-1, 1, 3, 5};
  EXPECT_TRUE(v == slice_vector(v2, 1, 3));
  std::vector<int> v3 = {3};
  EXPECT_TRUE(v3 == slice_vector(v, 1, 1));
}

TEST(outOfRangeTest, false) {
  EXPECT_EQ(binary_search(v, 0), false);
  EXPECT_EQ(binary_search(v, 6), false);
}

std::vector<int> v1;
for (int i = 0; i < 100; i++) {
  v1.push_back(i * 2);
}

TEST(included, true) {
  EXPECT_TRUE(binary_search(v, 1));
  EXPECT_TRUE(binary_search(v, 3));
  EXPECT_TRUE(binary_search(v, 3));
}
