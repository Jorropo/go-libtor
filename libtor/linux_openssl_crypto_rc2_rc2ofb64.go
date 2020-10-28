// go-libtor - Self-contained Tor from Go
// Copyright (c) 2018 Péter Szilágyi. All rights reserved.
// +build linux android
// +build staticOpenssl

package libtor

/*
#define DSO_NONE
#define OPENSSLDIR "/usr/local/ssl"
#define ENGINESDIR "/usr/local/lib/engines"

#include <../crypto/rc2/rc2ofb64.c>
*/
import "C"
