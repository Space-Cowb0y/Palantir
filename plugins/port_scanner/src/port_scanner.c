// port scanner
#include "../../../sdk/security_plugin.h"
static const st_host_api* HOST;

static int init(const st_host_api* host){ HOST = host; HOST->log_fn(1, "port_scanner init"); return 0; }
static int start(){ HOST->log_fn(1, "port_scanner start"); return 0; }
static int stop(int t){ (void)t; HOST->log_fn(1, "port_scanner stop"); return 0; }
static void shutdown(){ HOST->log_fn(1, "port_scanner bye"); }

int st_get_plugin_info(st_plugin_info* out){
  *out = (st_plugin_info){ "port_scanner", "0.1.0", 1, 0 };
  return 0;
}
int st_get_plugin_api(st_plugin_vtable* out){
  *out = (st_plugin_vtable){ init, start, stop, shutdown };
  return 0;
}
