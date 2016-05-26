// +build darwin

package CF

// #include <CoreFoundation/CFString.h>
import "C"
import (
	"reflect"
	"unsafe"
)

// The StringRef type is a reference to a Core Foundation string object.
//
// https://developer.apple.com/library/mac/documentation/CoreFoundation/Reference/CFStringRef/index.html#//apple_ref/c/tdef/CFStringRef
type StringRef unsafe.Pointer

// StringCreate takes a Go string as argument and creates a String object
// that represents the same content, then returns a reference to the newly
// created string.
//
// It is the program's responsibility to release the object returned by this
// function with a call to Release.
//
// https://developer.apple.com/library/mac/documentation/CoreFoundation/Reference/CFStringRef/index.html#//apple_ref/c/func/CFStringCreateWithBytes
func StringCreate(s string) StringRef {
	h := (*reflect.StringHeader)(unsafe.Pointer(&s))
	return StringRef(C.CFStringCreateWithBytes(
		nil,
		(*C.UInt8)(unsafe.Pointer(h.Data)),
		C.CFIndex(len(s)),
		C.kCFStringEncodingUTF8,
		0,
	))
}

// GoString creates a new Go string value with a content equivalent to the
// StringRef object passed as argument.
func GoString(s StringRef) string {
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
