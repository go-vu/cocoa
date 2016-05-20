// +build darwin

package cocoa

import (
	"image"
	"runtime"
)

// The Image type wraps and manages a CGImageRef value.
//
// The intent for Image values is to provide automatic memory management for
// native values from the Cocoa framework by leveraging the Go GC, basically
// the lifetime of the managed CGImageRef is bound to the lifetime of the
// Image instance.
//
// The Image type also offers an API that betters integrates with Go code by
// implementing standard interfaces and exposing methods.
type Image struct {
	ref CGImageRef
}

// NewImage creates a new Image value with a content equivalent to the Go
// image passed as argument.
func NewImage(img image.Image) *Image {
	return NewImageWrap(CGImageCreate(img))
}

// NewImageWarp creates a new Image value that wraps the CGImageRef value
// passed as argument.
//
// The returned Image value becomes the owner of the CGImageRef, it will
// automatically release it when it gets garbage collected.
func NewImageWrap(ref CGImageRef) *Image {
	img := &Image{ref}
	runtime.SetFinalizer(img, (*Image).release)
	return img
}

// Ref returns the CGImageRef wrapped by the Image value it is called on.
func (img *Image) Ref() CGImageRef {
	return img.ref
}

func (img *Image) release() {
	if ref := img.ref; ref != nil {
		img.ref = nil
		CFRelease(CFTypeRef(ref))
	}
}
