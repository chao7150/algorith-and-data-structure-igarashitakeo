#include <vector>
#include <tuple>

std::tuple<bool, int> binary_search(const std::vector<int> v, const int target, const int loop = 1);
std::vector<int> slice_vector(const std::vector<int> v, const int first,
                              const int last);