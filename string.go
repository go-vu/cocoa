// +build darwin

package cocoa

import "runtime"

// The String type wraps and manages a CFStringRef value.
//
// The intent for String values is to provide automatic memory management for
// native values from the Cocoa framework by leveraging the Go GC, basically
// the lifetime of the managed CFStringRef is bound to the lifetime of the
// String instance.
//
// The String type also offers an API that betters integrates with Go code by
// implementing standard interfaces and exposing methods.
type String struct {
	ref CFStringRef
}

// NewString creates a new String value with a content equivalent to the Go
// string passed as argument.
func NewString(s string) *String {
	return NewStringWrap(CFStringCreate(s))
}

// NewStringWarp creates a new String value that wraps the CFStringRef value
// passed as argument.
//
// The returned String value becomes the owner of the CFStringRef, it will
// automatically release it when it gets garbage collected.
func NewStringWrap(ref CFStringRef) *String {
	s := &String{ref}
	runtime.SetFinalizer(s, (*String).release)
	return s
}

// String returns the content of the String object as a Go string value.
func (s *String) String() string {
	return GoString(s.ref)
}

// Ref returns the CFStringRef wrapped by the String value it is called on.
func (s *String) Ref() CFStringRef {
	return s.ref
}

func (s *String) release() {
	if ref := s.ref; ref != nil {
		s.ref = nil
		CFRelease(CFTypeRef(ref))
	}
}
