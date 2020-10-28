// go-libtor - Self-contained Tor from Go
// Copyright (c) 2018 Péter Szilágyi. All rights reserved.
// +build linux android
// +build staticZlib

package libtor

/*
#include <../zlib/gzwrite.c>
*/
import "C"
