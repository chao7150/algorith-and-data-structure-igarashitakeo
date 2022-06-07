#include <gtest/gtest.h>

#include "hash_chain.h"

TEST(hash_chain, insert) {
  HashChain *h = new HashChain;
  insert(h, 115);
  EXPECT_EQ(h->data[5][0], 115);
  insert(h, 5);
  EXPECT_EQ(h->data[5][1], 5);
}

TEST(hash_chain, member) {
  HashChain *h = new HashChain;
  insert(h, 115);
  insert(h, 5);
  EXPECT_EQ(member(h, 5), true);
  EXPECT_EQ(member(h, 3), false);
}
