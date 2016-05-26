// +build darwin

package CG

import (
	"bufio"
	"image"
	"image/color"
	"image/draw"
	"os"
	"testing"

	_ "image/png"

	"github.com/go-vu/cocoa/CF"
)

func TestImageCreate(t *testing.T) {
	for _, test := range gopherImages {
		img := ImageCreate(test)
		CF.Retain(CF.TypeRef(img))
		CF.Release(CF.TypeRef(img))
		CF.Release(CF.TypeRef(img))
	}
}

func TestImageCreateNoCopy(t *testing.T) {
	for _, test := range gopherImages {
		img := ImageCreateNoCopy(test)
		CF.Retain(CF.TypeRef(img))
		CF.Release(CF.TypeRef(img))
		CF.Release(CF.TypeRef(img))
	}
}

func TestImageCreatePanic(t *testing.T) {
	defer func() { recover() }()
	ImageCreate(image.NewUniform(color.Black))
	t.Error("calling ImageCreate with an unsupported image type did not panic!")
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
