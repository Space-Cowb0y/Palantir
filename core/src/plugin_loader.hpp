// plugin loader header
#pragma once
#include <string>
#include <vector>
struct LoadedPlugin {
  void* handle;
  std::string name, version;
  int (*start)();
  int (*stop)(int);
  void (*shutdown)();
};
bool load_plugin(const std::string& path, LoadedPlugin& out);
