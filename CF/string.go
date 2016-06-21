// +build darwin

package CF

// #include <CoreFoundation/CFString.h>
import "C"
import (
	"fmt"
	"reflect"
	"unsafe"
)

// The StringRef type is a reference to a Core Foundation string object.
//
// https://developer.apple.com/library/mac/documentation/CoreFoundation/Reference/CFStringRef/index.html#//apple_ref/c/tdef/CFStringRef
type StringRef TypeRef

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
	return StringRef(unsafe.Pointer(C.CFStringCreateWithBytes(
		nil,
		(*C.UInt8)(unsafe.Pointer(h.Data)),
		C.CFIndex(len(s)),
		C.kCFStringEncodingUTF8,
		0,
	)))
}

// GoString creates a new Go string value with a content equivalent to the
// StringRef object passed as argument.
func GoString(s StringRef) string {
	ptr := C.CFStringGetCStringPtr(unsafe.Pointer(s), C.kCFStringEncodingUTF8)

	if ptr != nil {
		return C.GoString(ptr)
	}

	n := C.CFStringGetLength(unsafe.Pointer(s))
	b := make([]byte, int(4*n))
	C.CFStringGetBytes(
		unsafe.Pointer(s),
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

// Retain increases the refence counter of the Core Foundation string passed
// as argument.
//
// https://developer.apple.com/library/mac/documentation/CoreFoundation/Reference/CFTypeRef/#//apple_ref/c/func/CFRetain
func (s StringRef) Retain() {
	TypeRef(s).Retain()
}

// Release decreases the reference counter of the Core Foundation string
// passed as argument.
//
// https://developer.apple.com/library/mac/documentation/CoreFoundation/Reference/CFTypeRef/#//apple_ref/c/func/CFRelease
func (s StringRef) Release() {
	TypeRef(s).Release()
}

// Length returns the number of characters in the string it's called on.
//
// https://developer.apple.com/library/mac/documentation/CoreFoundation/Reference/CFStringRef/#//apple_ref/c/func/CFStringGetLength
func (s StringRef) Length() int {
	return int(C.CFStringGetLength(unsafe.Pointer(s)))
}

// String statisfies the fmt.Stringer interface.
func (s StringRef) String() string {
	return GoString(s)
}

// GoString satisfies the fmt.GoStringer interface.
func (s StringRef) GoString() string {
	return fmt.Sprintf("%v", s.String())
}
