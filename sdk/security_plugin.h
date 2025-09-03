// Palantir SDK ABI
#pragma once
#include <stdint.h>
#ifdef _WIN32
  #define ST_EXPORT __declspec(dllexport)
#else
  #define ST_EXPORT __attribute__((visibility("default")))
#endif

typedef struct {
  void (*log_fn)(int level, const char* msg);
  int  (*publish_event)(const uint8_t* data, uint32_t len);
  const char* (*get_config)(const char* key);
} st_host_api;

typedef struct {
  const char* name;
  const char* version;
  uint32_t abi_version;
  uint32_t capabilities;
} st_plugin_info;

typedef struct {
  int  (*init)(const st_host_api* host);
  int  (*start)();
  int  (*stop)(int timeout_ms);
  void (*shutdown)();
} st_plugin_vtable;

#ifdef __cplusplus
extern "C" {
#endif
ST_EXPORT int st_get_plugin_info(st_plugin_info* out);
ST_EXPORT int st_get_plugin_api(st_plugin_vtable* out);
#ifdef __cplusplus
}
#endif
