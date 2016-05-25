// +build darwin

package cocoa

import (
	"bufio"
	"image"
	"image/draw"
	"os"
	"testing"
)

func TestNewImage(t *testing.T) {
	for _, test := range gopherImages {
		img := NewImage(test)

		if img.Ref() == nil {
			t.Error("invalid nil image reference")
		}
	}
}

func TestNewImageWrap(t *testing.T) {
	for _, test := range gopherImages {
		img1 := CGImageCreate(test)
		img2 := NewImageWrap(img1)

		if ref := img2.Ref(); ref != img1 {
			t.Error("invalid image reference:", ref)
		}
	}
}

func TestImageRelease(t *testing.T) {
	for _, test := range gopherImages {
		img := NewImage(gopherNRGBA)
		img.release()

		if img.Ref() != nil {
			t.Error("invalid non-nil image reference found after releasing")
		}
	}
}

func loadImage() image.Image {
	f, _ := os.Open("fixtures/gopher.png")
	i, _, _ := image.Decode(bufio.NewReader(f))
	f.Close()
	return i
}

func copyImage(dst draw.Image, src image.Image) image.Image {
	draw.Draw(dst, dst.Bounds(), src, image.Point{}, draw.Over)
	return dst
}

func imageToAlpha(img image.Image) image.Image {
	return copyImage(image.NewAlpha(img.Bounds()), img)
}

func imageToAlpha16(img image.Image) image.Image {
	return copyImage(image.NewAlpha16(img.Bounds()), img)
}

func imageToGray(img image.Image) image.Image {
	return copyImage(image.NewGray(img.Bounds()), img)
}

func imageToGray16(img image.Image) image.Image {
	return copyImage(image.NewGray16(img.Bounds()), img)
}

func imageToRGBA(img image.Image) image.Image {
	return copyImage(image.NewRGBA(img.Bounds()), img)
}

func imageToRGBA64(img image.Image) image.Image {
	return copyImage(image.NewRGBA64(img.Bounds()), img)
}

func imageToNRGBA64(img image.Image) image.Image {
	return copyImage(image.NewNRGBA64(img.Bounds()), img)
}

func imageToCMYK(img image.Image) image.Image {
	return copyImage(image.NewCMYK(img.Bounds()), img)
}

var (
	gopherNRGBA   = loadImage()
	gopherAlpha   = imageToAlpha(gopherNRGBA)
	gopherAlpha16 = imageToAlpha16(gopherNRGBA)
	gopherGray    = imageToGray(gopherNRGBA)
	gopherGray16  = imageToGray16(gopherNRGBA)
	gopherRGBA    = imageToRGBA(gopherNRGBA)
	gopherRGBA64  = imageToRGBA64(gopherNRGBA)
	gopherNRGBA64 = imageToNRGBA64(gopherNRGBA)
	gopherCMYK    = imageToCMYK(gopherNRGBA)

	gopherImages = [...]image.Image{
		gopherNRGBA,
		gopherAlpha,
		gopherAlpha16,
		gopherGray,
		gopherGray16,
		gopherRGBA,
		gopherRGBA64,
		gopherNRGBA64,
		gopherCMYK,
	}
)
