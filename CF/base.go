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
type TypeRef uintptr

// The TypeID type is used to provide a unique identifier to the type of Core
// Foundation object.
//
// https://developer.apple.com/library/mac/documentation/CoreFoundation/Reference/CFTypeRef/#//apple_ref/c/tdef/CFTypeID
type TypeID uint64

// GetTypeID returns the TypeID representing the type of the Core Foundation
// object passed as argument.
//
// https://developer.apple.com/library/mac/documentation/CoreFoundation/Reference/CFTypeRef/#//apple_ref/c/func/CFGetTypeID
func (obj TypeRef) GetTypeID() TypeID {
	return TypeID(C.CFGetTypeID(C.CFTypeRef(obj)))
}

// Retain increases the refence counter of the Core Foundation object passed
// as argument.
//
// https://developer.apple.com/library/mac/documentation/CoreFoundation/Reference/CFTypeRef/#//apple_ref/c/func/CFRetain
func (obj TypeRef) Retain() {
	C.CFRetain(C.CFTypeRef(obj))
}

// Release decreases the reference counter of the Core Foundation object
// passed as argument.
//
// https://developer.apple.com/library/mac/documentation/CoreFoundation/Reference/CFTypeRef/#//apple_ref/c/func/CFRelease
func (obj TypeRef) Release() {
	C.CFRelease(C.CFTypeRef(obj))
}

// CopyDescription returns a string representation of the object passed as
// argument.
//
// https://developer.apple.com/library/mac/documentation/CoreFoundation/Reference/CFTypeRef/#//apple_ref/c/func/CFCopyDescription
func (obj TypeRef) CopyDescription() StringRef {
	return StringRef(unsafe.Pointer(C.CFCopyDescription(C.CFTypeRef(obj))))
}

// String satisfies the fmt.Stringer interface.
func (obj TypeRef) String() string {
	s := obj.CopyDescription()
	defer s.Release()
	return GoString(s)
}

// CopyTypeIDDescription returns a string representation of the type id passed
// as argument.
//
// https://developer.apple.com/library/mac/documentation/CoreFoundation/Reference/CFTypeRef/#//apple_ref/c/func/CFCopyTypeIDDescription
func (id TypeID) CopyTypeIDDescription() StringRef {
	return StringRef(unsafe.Pointer(C.CFCopyTypeIDDescription(C.CFTypeID(id))))
}

// String satisfies the fmt.Stringer interface.
func (id TypeID) String() string {
	s := id.CopyTypeIDDescription()
	defer s.Release()
	return GoString(s)
}

// Equal tests two object for equality.
//
// https://developer.apple.com/library/mac/documentation/CoreFoundation/Reference/CFTypeRef/#//apple_ref/c/func/CFEqual
func Equal(obj1 TypeRef, obj2 TypeRef) bool {
	return C.CFEqual(C.CFTypeRef(obj1), C.CFTypeRef(obj2)) != 0
}
