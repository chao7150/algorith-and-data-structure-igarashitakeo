int euclidean(const int a, const int b) {
  int big;
  int small;
  if (a > b) {
    big = a;
    small = b;
  } else {
    big = b;
    small = a;
  }
  int remainder = big % small;
  if (remainder == 0) {
    return small;
  }
  return euclidean(small, remainder);
}