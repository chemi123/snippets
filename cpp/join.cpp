#include <iostream>
#include <iterator>
#include <sstream>
#include <vector>

using namespace std;

int main() {
  vector<string> strVec{"a", "b", "c"};

  ostringstream oss;
  copy(strVec.begin(), strVec.end(), ostream_iterator<string>(oss, ","));
  string str = oss.str();
  str.erase(str.size() - 1);
  cout << str << endl;
  return 0;
}
