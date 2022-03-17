#include <iostream>
#include <tuple>
#include <vector>

std::vector<int> slice_vector(const std::vector<int> v, const int first,
                              const int last) {
  std::vector<int> rv(v.begin() + first, v.begin() + last + 1);
  return rv;
};

std::tuple<bool, int> binary_search(const std::vector<int> v, const int target,
                                    const int loop) {
  int length = v.size();
  if (length == 0)
    return {false, loop};
  if (target < v.front() || v.back() < target)
    return {false, loop};
  // 小数点以下切り捨て
  int mid_point = length / 2;
  if (v[mid_point] == target)
    return {true, loop};
  if (v[mid_point] > target) {
    return binary_search(slice_vector(v, 0, mid_point - 1), target, loop + 1);
  }
  
  return binary_search(slice_vector(v, mid_point + 1, length - 1), target,
                       loop + 1);
};
