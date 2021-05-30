#include <iostream>
#include <vector>

using namespace std;

// vectorのiteratorから要素番号を求める方法

int main() {
  vector<int> v{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10};
  auto it = lower_bound(v.begin(), v.end(), 7);
  cout << distance(v.begin(), it) << endl;
  return 0;
}
