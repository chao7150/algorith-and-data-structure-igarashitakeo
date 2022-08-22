#include <gtest/gtest.h>

#include "hash_open_address.h"

TEST(hash_open_address, member) {
  HashOpenAddress *h = new HashOpenAddress;
  EXPECT_EQ(member(h, 5), false);
  insert(h, 5);
  EXPECT_EQ(member(h, 5), true);
  EXPECT_EQ(member(h, 15), false);
  insert(h, 15);
  EXPECT_EQ(member(h, 15), true);
}

TEST(hash_open_address, insert) {
  HashOpenAddress *h = new HashOpenAddress;
  insert(h, 5);
  insert(h, 15);
  insert(h, 25);
  EXPECT_EQ(h->data[5], 5);
  EXPECT_EQ(h->data[6], 15);
  EXPECT_EQ(h->data[7], 25);
}
