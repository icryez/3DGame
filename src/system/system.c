#include "system.h"
#include <stdio.h>
#include <wayland-client.h>

int connectToWayland() {
  struct wl_display *display = wl_display_connect(NULL);
  if (!display) {
    fprintf(stderr, "Failed to connect to Wayland display.\n");
    return 1;
  }
  fprintf(stderr, "Connection established!\n");

  wl_display_disconnect(display);
  return 0;
}
