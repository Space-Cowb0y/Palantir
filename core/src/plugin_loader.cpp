// plugin loader
#include "plugin_loader.hpp"
#include <iostream>
#include <dlfcn.h> // use <windows.h> e ajuste no Windows
#include "../../sdk/security_plugin.h"

static void log_fn(int level, const char* msg){ std::cerr << "[P" << level << "] " << msg << "\n"; }
static int publish_event(const uint8_t* data, uint32_t len){ (void)len; /* TODO: enviar p/ Agent API */ return 0; }
static const char* get_config(const char* key){ (void)key; return ""; }

bool load_plugin(const std::string& path, LoadedPlugin& out){
  void* h = dlopen(path.c_str(), RTLD_NOW);
  if(!h){ std::cerr << "dlopen fail\n"; return false; }
  auto info = (int(*)(st_plugin_info*)) dlsym(h, "st_get_plugin_info");
  auto api  = (int(*)(st_plugin_vtable*)) dlsym(h, "st_get_plugin_api");
  st_plugin_info pi; st_plugin_vtable vt;
  if(info(&pi)!=0 || api(&vt)!=0) return false;
  st_host_api host{ &log_fn, &publish_event, &get_config };
  vt.init(&host);
  out = { h, pi.name, pi.version, vt.start, vt.stop, vt.shutdown };
  return true;
}
