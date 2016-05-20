// +build darwin

package cocoa

import (
	"testing"

	_ "image/png"
)

func TestCGImageCreate(t *testing.T) {
	img := CGImageCreate(gopher)
	CFRetain(CFTypeRef(img))
	CFRelease(CFTypeRef(img))
	CFRelease(CFTypeRef(img))
}

func TestCGImageCreateNoCopy(t *testing.T) {
	img := CGImageCreateNoCopy(gopher)
	CFRetain(CFTypeRef(img))
	CFRelease(CFTypeRef(img))
	CFRelease(CFTypeRef(img))
}
