// +build darwin

// Package cocoa provides functions and types that are shared between go-vu
// drivers for Apple platforms.
package cocoa

// #cgo CFLAGS: -x objective-c -fobjc-arc -Wno-unused-parameter
// #cgo LDFLAGS: -framework AppKit
import "C"
