// +build darwin

package cocoa

// #include <CoreFoundation/CFString.h>
import "C"
import (
	"reflect"
	"unsafe"
)

// The CFTypeRef type is an untyped reference to any Core Foundation object.
//
// https://developer.apple.com/library/mac/documentation/CoreFoundation/Reference/CFTypeRef/#//apple_ref/c/tdef/CFTypeRef
type CFTypeRef unsafe.Pointer

// CFRetain increases the refence counter of the Core Foundation object passed
// as argument.
//
// https://developer.apple.com/library/mac/documentation/CoreFoundation/Reference/CFTypeRef/#//apple_ref/c/func/CFRetain
func CFRetain(obj CFTypeRef) {
	C.CFRetain(C.CFTypeRef(obj))
}

// CFRelease decreases the reference counter of the Core Foundation object
// passed as argument.
//
// https://developer.apple.com/library/mac/documentation/CoreFoundation/Reference/CFTypeRef/#//apple_ref/c/func/CFRelease
func CFRelease(obj CFTypeRef) {
	C.CFRelease(C.CFTypeRef(obj))
}

// The CFStringRef type is a reference to a Core Foundation string object.
//
// https://developer.apple.com/library/ios/documentation/CoreFoundation/Reference/CFStringRef/index.html#//apple_ref/c/tdef/CFStringRef
type CFStringRef C.CFStringRef

// CFStringCreate takes a Go string as argument and creates a CFString object
// that represents the same content, then returns a reference to the newly
// created string.
//
// It is the program's responsibility to release the object returned by this
// function with a call to CFStringRelease or CFRelease.
//
// https://developer.apple.com/library/ios/documentation/CoreFoundation/Reference/CFStringRef/index.html#//apple_ref/c/func/CFStringCreateWithBytes
func CFStringCreate(s string) CFStringRef {
	h := (*reflect.StringHeader)(unsafe.Pointer(&s))
	return CFStringRef(C.CFStringCreateWithBytes(
		nil,
		(*C.UInt8)(unsafe.Pointer(h.Data)),
		C.CFIndex(len(s)),
		C.kCFStringEncodingUTF8,
		0,
	))
}

// CFStringRetain increases the reference counter of the string object passed as
// argument.
//
// https://developer.apple.com/library/mac/documentation/CoreFoundation/Reference/CFTypeRef/#//apple_ref/c/func/CFRetain
func CFStringRetain(s CFStringRef) {
	CFRetain(CFTypeRef(s))
}

// CFStringRelease decreases the reference counter of the string object passed
// as argument.
//
// https://developer.apple.com/library/mac/documentation/CoreFoundation/Reference/CFTypeRef/#//apple_ref/c/func/CFRelease
func CFStringRelease(s CFStringRef) {
	CFRelease(CFTypeRef(s))
}

// GoString creates a new Go string value with a content equivalent to the
// CFStringRef object passed as argument.
func GoString(s CFStringRef) string {
	ptr := C.CFStringGetCStringPtr(s, C.kCFStringEncodingUTF8)

	if ptr != nil {
		return C.GoString(ptr)
	}

	n := C.CFStringGetLength(s)
	b := make([]byte, int(4*n))
	C.CFStringGetBytes(
		s,
		C.CFRange{0, n},
		C.kCFStringEncodingUTF8,
		'?',
		0,
		(*C.UInt8)(unsafe.Pointer(&b[0])),
		C.CFIndex(len(b)),
		&n,
	)

	return string(b[:n])
}
