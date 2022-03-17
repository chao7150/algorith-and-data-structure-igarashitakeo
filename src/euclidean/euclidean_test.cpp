#include <gtest/gtest.h>

#include "euclidean.h"

TEST(euclidean, gcm) {
  EXPECT_EQ(euclidean(3, 6), 3);
  EXPECT_EQ(euclidean(4, 6), 2);
  EXPECT_EQ(euclidean(3355, 2379), 61);
  EXPECT_EQ(euclidean(1, 1), 1);
}