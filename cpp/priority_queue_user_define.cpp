#include <iostream>
#include <queue>
#include <vector>

using namespace std;

struct Node {
  int _num;
  Node(int num) : _num(num) {}
};

struct CompNode {
  bool operator()(const Node& n1, const Node& n2) {
    return n1._num < n2._num;
  }
};

int main() {
  priority_queue<Node, vector<Node>, CompNode> q;
  q.push(Node(6));
  q.push(Node(9));
  q.push(Node(4));
  q.push(Node(10));
  q.push(Node(1));
  q.push(Node(2));
  q.push(Node(8));
  q.push(Node(7));
  q.push(Node(5));
  q.push(Node(3));
  while (!q.empty()) {
    cout << q.top()._num << endl;
    q.pop();
  }
  return 0;
}
