// +build darwin

package CF

// #cgo CFLAGS: -Wno-unused-parameter
// #cgo LDFLAGS: -framework CoreFoundation
//
// #include <CoreFoundation/CFBase.h>
import "C"
import "unsafe"

// The TypeRef type is an untyped reference to any Core Foundation object.
//
// https://developer.apple.com/library/mac/documentation/CoreFoundation/Reference/CFTypeRef/#//apple_ref/c/tdef/CFTypeRef
type TypeRef unsafe.Pointer

// Retain increases the refence counter of the Core Foundation object passed
// as argument.
//
// https://developer.apple.com/library/mac/documentation/CoreFoundation/Reference/CFTypeRef/#//apple_ref/c/func/CFRetain
func Retain(obj TypeRef) {
	C.CFRetain(C.CFTypeRef(obj))
}

// Release decreases the reference counter of the Core Foundation object
// passed as argument.
//
// https://developer.apple.com/library/mac/documentation/CoreFoundation/Reference/CFTypeRef/#//apple_ref/c/func/CFRelease
func Release(obj TypeRef) {
	C.CFRelease(C.CFTypeRef(obj))
}
