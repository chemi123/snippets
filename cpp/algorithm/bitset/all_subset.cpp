#include <iostream>
#include <vector>

using namespace std;

int main() {
  int n = 5;
  for (int bit = 0; bit < (1 << n); ++bit) {
    vector<int> vec;
    for (int i = 0; i < n; ++i) {
      if (bit & (1 << i)) {
        vec.emplace_back(i);
      }
    }

    if (vec.size() == (size_t)0) continue;
    for (auto& e : vec) {
      cout << e << " ";
    }
    cout << endl;
  }
  return 0;
}