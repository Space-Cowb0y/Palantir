// Palantir core main
#include "plugin_loader.hpp"
#include <iostream>

int main(){
  LoadedPlugin p;
  if(load_plugin("./build/plugins/port_scanner/libport_scanner.so", p)){
    std::cout << "Loaded: " << p.name << " v" << p.version << "\n";
    p.start();
    p.stop(500);
    p.shutdown();
  }
  return 0;
}
