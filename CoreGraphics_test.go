// +build darwin

package cocoa

import (
	"bufio"
	"image"
	"os"
	"testing"

	_ "image/png"
)

func TestCGImageCreate(t *testing.T) {
	img := CGImageCreate(gopher)
	CGImageRetain(img)
	CGImageRelease(img)
	CGImageRelease(img)
}

func TestCGImageCreateNoCopy(t *testing.T) {
	img := CGImageCreateNoCopy(gopher)
	CGImageRetain(img)
	CGImageRelease(img)
	CGImageRelease(img)
}

var gopher image.Image

func init() {
	f, _ := os.Open("fixtures/gopher.png")
	gopher, _, _ = image.Decode(bufio.NewReader(f))
	f.Close()
}
