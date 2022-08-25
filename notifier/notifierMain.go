package notifier

// #cgo LDFLAGS: -framework CoreFoundation -framework IOKit
// int CanSleep();
// void WillWake();
// void WillSleep();
// #include "main.h"
import "C"

func StartNotifier() {
	C.registerNotifications()
}

func StopNotifier() {
	C.unregisterNotifications()
}
