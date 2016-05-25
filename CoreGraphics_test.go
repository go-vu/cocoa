// +build darwin

package cocoa

import (
	"image"
	"testing"

	"image/color"
	_ "image/png"
)

func TestCGImageCreate(t *testing.T) {
	for _, test := range gopherImages {
		img := CGImageCreate(test)
		CFRetain(CFTypeRef(img))
		CFRelease(CFTypeRef(img))
		CFRelease(CFTypeRef(img))
	}
}

func TestCGImageCreateNoCopy(t *testing.T) {
	for _, test := range gopherImages {
		img := CGImageCreateNoCopy(test)
		CFRetain(CFTypeRef(img))
		CFRelease(CFTypeRef(img))
		CFRelease(CFTypeRef(img))
	}
}

func TestCGImageCreatePanic(t *testing.T) {
	defer func() { recover() }()
	CGImageCreate(image.NewUniform(color.Black))
	t.Error("calling CGImageCreate with an unsupported image type did not panic!")
}
