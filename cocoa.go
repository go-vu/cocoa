// +build darwin

// Package cocoa provides functions and types that are shared between go-vu
// drivers for Apple platforms.
//
// The intent of thie package is not to provide a full interface to the Cocoa
// framework in go but rather to share reusable code and expose convenient
// abstractions in the context of building OSX and iOS drivers for the go-vu
// project.
package cocoa

// #cgo CFLAGS: -x objective-c -fobjc-arc -Wno-unused-parameter
// #cgo LDFLAGS: -framework AppKit
import "C"
