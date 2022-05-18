#include <gtest/gtest.h>

#include "binary_search_tree.h"

TEST(binary_search_tree, insert) {
  Node *n = new Node{10};
  insert(n, 9);
  EXPECT_EQ(n->left->value, 9);
  insert(n, 12);
  EXPECT_EQ(n->right->value, 12);
}

TEST(binary_search_tree, search) {
  Node *n = new Node{10};
  insert(n, 9);
  insert(n, 12);
  insert(n, 11);
  EXPECT_EQ(search(n, 11), true);
  EXPECT_EQ(search(n, 13), false);
}

TEST(binary_search_tree, remove_zero_child) {
  Node *n = new Node{5};
  insert(n, 4);
  EXPECT_EQ(n->left->value, 4);
  remove(n, 4);
  EXPECT_EQ(n->left, nullptr);
}

TEST(binary_search_tree, remove) {
  Node *n = new Node{5};
  insert(n, 3);
  insert(n, 2);
  insert(n, 4);
  insert(n, 9);
  insert(n, 8);
  insert(n, 11);
  insert(n, 10);
  insert(n, 13);
  remove(n, 9);
  EXPECT_EQ(n->right->value, 10);
  EXPECT_EQ(n->right->right->value, 11);
  EXPECT_EQ(n->right->right->left, nullptr);
  EXPECT_EQ(n->right->right->right->value, 13);
}