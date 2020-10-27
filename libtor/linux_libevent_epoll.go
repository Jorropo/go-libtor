// go-libtor - Self-contained Tor from Go
// Copyright (c) 2018 Péter Szilágyi. All rights reserved.
// +build linux android
// +build staticLibevent

package libtor

/*
#ifdef PREFIX_GO_LIBTOR
# pragma extern_prefix GO_LIBTOR_
#endif

#include <compat/sys/queue.h>
#include <../epoll.c>
*/
import "C"
