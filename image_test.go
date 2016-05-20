// +build darwin

package cocoa

import (
	"bufio"
	"image"
	"os"
	"testing"
)

var gopher image.Image

func init() {
	f, _ := os.Open("fixtures/gopher.png")
	gopher, _, _ = image.Decode(bufio.NewReader(f))
	f.Close()
}

func TestNewImage(t *testing.T) {
	img := NewImage(gopher)

	if img.Ref() == nil {
		t.Error("invalid nil image reference")
	}
}

func TestNewImageWrap(t *testing.T) {
	img1 := CGImageCreate(gopher)
	img2 := NewImageWrap(img1)

	if ref := img2.Ref(); ref != img1 {
		t.Error("invalid image reference:", ref)
	}
}

func TestImageRelease(t *testing.T) {
	img := NewImage(gopher)
	img.release()

	if img.Ref() != nil {
		t.Error("invalid non-nil image reference found after releasing")
	}
}
