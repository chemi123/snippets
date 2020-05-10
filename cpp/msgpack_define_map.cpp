#include <msgpack.hpp>
#include <vector>
#include <unordered_map>
#include <iostream>

using namespace std;

struct Segment {
  int num;
  string str;
  Segment() : num(0), str("") {}
  Segment(int num, string& str) : num(num), str(str) {}

  MSGPACK_DEFINE_MAP(num, str);
};

struct Model {
  vector<Segment> segments;

  Model() : segments() {}
  Model(vector<Segment>& segments) : segments(segments) {}
  MSGPACK_DEFINE_MAP(segments);
};

typedef unordered_map<string, Model> ModelMap;

int main() {
  vector<Segment> segments;
  Segment s1, s2;
  s1.num = 1;
  s1.str = "segment1";
  s2.num = 2;
  s2.str = "segment2";
  segments.emplace_back(s1);
  segments.emplace_back(s2);
  Model m(segments);
  ModelMap modelMap = {{"model1", m}};

  msgpack::sbuffer sbuf;
  msgpack::pack(sbuf, modelMap);

  msgpack::object_handle oh = msgpack::unpack(sbuf.data(), sbuf.size());
  msgpack::object obj = oh.get();

  ModelMap modelMap2;
  obj.convert(modelMap2);
  for (const auto& [key, model] : modelMap2) {
    cout << key << endl;
    for (const auto& segment: model.segments) {
      cout << "str: " << segment.str << endl;
      cout << "num: " << segment.num << endl;
    }
  }

  return 0;
}
