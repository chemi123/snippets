#include <iostream>
#include <cmath>

using namespace std;

// 桁数を求める

int main() {
  int num = 1;
  cout << num << ": " << floor(log10(num)) + 1 << endl;
  num = 10;
  cout << num << ": " << floor(log10(num)) + 1 << endl;
  num = 100;
  cout << num << ": " << floor(log10(num)) + 1 << endl;
  num = 1000;
  cout << num << ": " << floor(log10(num)) + 1 << endl;
  num = 20438;
  cout << num << ": " << floor(log10(num)) + 1 << endl;
  return 0;
}

