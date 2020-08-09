#include <iostream>
#include <vector>

using namespace std;

/*
  以下よりunion findを試してみる。acが取れていることを確認している。
  https://atcoder.jp/contests/atc001/tasks/unionfind_a
*/

class UnionFind {
public:
  UnionFind(int n) : _parent(std::vector<int>(n, -1)) {
  }

  int root(int x) {
    if (_parent[x] < 0) {
      return x;
    }
    _parent[x] = root(_parent[x]);
    return _parent[x];
  }

  void unite(int x, int y) {
    x = root(x);
    y = root(y);

    if (x == y) {
      return;
    }

    if (_parent[x] > _parent[y]) {
      swap(x, y);
    }

    _parent[x] += _parent[y];
    _parent[y] = x;
  }

  bool isSame(int x, int y) {
    return root(x) == root(y);
  }

  int size(int x) {
    return -_parent[root(x)];
  }

  int groups() {
    int res = 0;
    // 要素が0から始まるか1から始まるかによって初期値が0か1かで変わる
    for (size_t i = 1; i < _parent.size(); ++i) {
      if (_parent.at(i) < 0) {
        ++res;
      }
    }
    return res;
  }

private:
  vector<int> _parent;
};

int main() {
  int n, q;
  cin >> n >> q;
  UnionFind uf(n);
  for (int i = 0; i < q; ++i) {
    int query, x, y;
    cin >> query >> x >> y;
    if (query == 0) {
      uf.unite(x, y);
    } else {
      cout << (uf.isSame(x, y) ? "Yes" : "No") << endl;
    }
  }
  return 0;
}