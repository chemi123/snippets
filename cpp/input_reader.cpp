#include <iostream>

using namespace std;

/*
  以下の形式のデータ入力を読み取る例
  最初の行はデータセットの数、以下3行はそれぞれのデータセット
  3 
  1 2 3
  4 5 6
  7 8 9 
*/

int main() {
  int n;
  cin >> n;
  for (int i = 0; i < n; ++i) {
    int a, b, c;
    cin >> a >> b >> c;
    cout << a  << ", " << b << ", " << c << endl;
  }
  return 0;
}