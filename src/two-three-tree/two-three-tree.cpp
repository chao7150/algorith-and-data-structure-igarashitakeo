#include "two-three-tree.h"
#include <iostream>
#include <utility>

int min(Node *node) {
  if (node->value != -1) {
    return node->value;
  }
  return min(node->children[0]);
}

/**
 * rootノードだけを扱う
 */
Node *insert(Node *node, const int value) {
  // ただ1つしかノードがないとき、いきなりleafに到達する
  if (node->value != -1) {
    Node *leaf_node = new Node{value, {-1, -1}, {nullptr, nullptr, nullptr}};
    Node *new_root =
        new Node{-1, {leaf_node->value, -1}, {node, leaf_node, nullptr}};
    return new_root;
  }
  auto [left, right] = _insert(node, value);
  if (right == nullptr) {
    return left;
  }
  return new Node{-1, {min(right), -1}, {left, right, nullptr}};
}

std::pair<Node *, Node *> _insert(Node *node, const int value) {
  // leafの1つ上である
  if (node->children[0]->value != -1) {
    // 子は2つである
    if (node->children[2] == nullptr) {
      // 右に挿入する
      if (node->markers[0] < value) {
        node->children[2] =
            new Node{value, {-1, -1}, {nullptr, nullptr, nullptr}};
        node->markers[1] = node->children[2]->value;
        return {node, nullptr};
      } else if (node->children[0]->value < value) { // 真ん中に挿入する
        node->children[2] = node->children[1];
        node->children[1] =
            new Node{value, {-1, -1}, {nullptr, nullptr, nullptr}};
        node->markers[0] = node->children[1]->value;
        node->markers[1] = node->children[2]->value;
        return {node, nullptr};
      } else { // 左に挿入する
        node->children[2] = node->children[1];
        node->children[1] = node->children[0];
        node->children[0] =
            new Node{value, {-1, -1}, {nullptr, nullptr, nullptr}};
        node->markers[0] = node->children[1]->value;
        node->markers[1] = node->children[2]->value;
        return {node, nullptr};
      }
    } else { // 子は3つである
      // 入れる値が1番大きい
      if (node->markers[1] < value) {
        return {
            new Node{-1,
                     {node->children[1]->value, -1},
                     {node->children[0], node->children[1]}},
            new Node{-1,
                     {value, -1},
                     {node->children[2],
                      new Node{value, {-1, -1}, {nullptr, nullptr, nullptr}}}}};
      } else if (node->markers[0] < value) { // 入れる値が上から2番目
        return {
            new Node{-1,
                     {node->children[1]->value, -1},
                     {node->children[0], node->children[1]}},
            new Node{-1,
                     {node->children[2]->value, -1},
                     {new Node{value, {-1, -1}, {nullptr, nullptr, nullptr}},
                      node->children[2]}}};
      } else if (node->children[0]->value < value) { // 入れる値が上から3番目
        return {
            new Node{-1,
                     {value, -1},
                     {node->children[0],
                      new Node{value, {-1, -1}, {nullptr, nullptr, nullptr}}}},
            new Node{-1,
                     {node->children[2]->value, -1},
                     {node->children[1], node->children[2]}}};
      } else { // 入れる値が上から4番目
        return {
            new Node{-1,
                     {node->children[0]->value, -1},
                     {new Node{value, {-1, -1}, {nullptr, nullptr, nullptr}},
                      node->children[0]}},
            new Node{-1,
                     {node->children[2]->value, -1},
                     {node->children[1], node->children[2]}}};
      }
    }
  }
  // leafの1個上ではない
  // 子が2つである
  // 子が4つ以上になることはないので自身は分裂しない
  if (node->children[2] == nullptr) {
    // 中央部分木を更新する
    if (node->markers[0] < value) {
      auto [left, right] = _insert(node->children[1], value);
      node->children[1] = left;
      node->children[2] = right;
      if (right != nullptr)
        node->markers[1] = min(right);
      return {node, nullptr};
    } else { // 左部分木を更新する
      auto [left, right] = _insert(node->children[2], value);
      node->children[0] = left;
      if (right == nullptr) {
        return {node, nullptr};
      } else {
        node->children[2] = node->children[1];
        node->markers[1] = min(node->children[2]);
        node->children[1] = right;
        node->markers[0] = min(right);
        return {node, nullptr};
      }
    }
  } else {                          // 子が3つである
    if (value < node->markers[0]) { // 左部分木を更新する
      auto [left, right] = _insert(node->children[0], value);
      if (right == nullptr) {
        node->children[0] = left;
        return {node, nullptr};
      } else {
        return {new Node{-1, {min(right), -1}, {left, right}},
                new Node{-1,
                         {min(node->children[2]), -1},
                         {node->children[1], node->children[2]}}};
      }
    } else if (value < node->markers[1]) { // 中央部分木を更新する
      auto [left, right] = _insert(node->children[1], value);
      if (right == nullptr) {
        node->children[1] = left;
        node->markers[0] = min(left);
        return {node, nullptr};
      } else {
        return {new Node{-1, {min(left), -1}, {node->children[0], left}},
                new Node{-1,
                         {min(node->children[2]), -1},
                         {right, node->children[2]}}};
      }
    } else { // 右部分木を更新する
      auto [left, right] = _insert(node->children[2], value);
      if (right == nullptr) {
        node->children[2] = left;
        node->markers[1] = min(left);
        return {node, nullptr};
      } else {
        return {new Node{-1,
                         {min(node->children[1]), -1},
                         {node->children[0], node->children[1]}},
                new Node{-1, {min(right), -1}, {left, right}}};
      }
    }
  }
}
