#include <iostream>

using namespace std;

/*
  以下よりunion findを試してみる。acが取れていることを確認している。
  https://atcoder.jp/contests/atc001/tasks/unionfind_a
*/

int find(int* root, int x) {
  if (root[x] == x) {
    return x;
  }
  return root[x] = find(root, root[x]);
}

bool same(int* root, int x, int y) {
  return find(root, x) == find(root, y);
}

void unite(int* root, int x, int y) {
  find(root, x);
  root[x] = find(root, y);
}

int main() {
  int n, q;
  cin >> n >> q;
  int root[n+1];
  for (int i = 0; i <= n; ++i) {
    root[i] = i;
  }
  for (int i = 0; i < q; ++i) {
    int judge, x, y;
    cin >> judge >> x >> y;
    if (judge) {
      cout << (same(root, x, y) ? "Yes" : "No") << endl;
    } else {
      unite(root, x, y);
    }
  }
  return 0;
}