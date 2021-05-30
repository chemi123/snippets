#include <iostream>
#include <unordered_set>

using namespace std;

// mapやsetでpairをキーとするやり方

struct phash{
  inline size_t operator()(const pair<int,int> & p) const{
    const auto h1 = hash<int>()(p.first);
    const auto h2 = hash<int>()(p.second);
    return h1 ^ (h2 << 1);
  }
};

int main() {
  unordered_set<pair<int, int>, phash> s;
  s.emplace(make_pair(10, 20));
  auto itr = s.find(make_pair(10, 20));
  if (itr != s.end()) {
    cout << itr->first << " " << itr->second << endl;
  }
  return 0;
}
