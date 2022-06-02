#include <gtest/gtest.h>

#include "two-three-tree.h"

TEST(twothreetree, min) {}

TEST(twothreetree, popSingle) {
  Node *n = new Node{10, {-1, -1}, {nullptr, nullptr, nullptr}};
  EXPECT_EQ(n->value, 10);
  Node *n2 = insert(n, 20);
  EXPECT_EQ(n2->markers[0], 20);
  EXPECT_EQ(n2->children[0]->value, 10);
  EXPECT_EQ(n2->children[1]->value, 20);
  Node *n3 = insert(n2, 30);
  EXPECT_EQ(n3->markers[0], 20);
  EXPECT_EQ(n3->markers[1], 30);
  EXPECT_EQ(n3->children[0]->value, 10);
  EXPECT_EQ(n3->children[1]->value, 20);
  EXPECT_EQ(n3->children[2]->value, 30);
  Node *n4 = insert(n3, 40);
  EXPECT_EQ(n4->markers[0], 30);
  EXPECT_EQ(n4->children[0]->markers[0], 20);
  EXPECT_EQ(n4->children[0]->children[0]->value, 10);
  EXPECT_EQ(n4->children[0]->children[1]->value, 20);
  EXPECT_EQ(n4->children[1]->markers[0], 40);
  EXPECT_EQ(n4->children[1]->children[0]->value, 30);
  EXPECT_EQ(n4->children[1]->children[1]->value, 40);
  Node* n5 = insert(n4, 50);
  EXPECT_EQ(n5->children[1]->children[2]->value, 50);
  Node* n6 = insert(n5, 60);
  EXPECT_EQ(n6->markers[1], 50);
  Node* n7 = insert(n6, 35);
  EXPECT_EQ(n7->children[1]->markers[0], 35);
  EXPECT_EQ(n7->children[1]->markers[1], 40);
}

// 子ノードが末端ではないときの3番目入れ