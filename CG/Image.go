// +build darwin

package CG

// #include <CoreGraphics/CGImage.h>
import "C"
import (
	"fmt"
	"image"
	"unsafe"
)

// The ImageRef type is a reference to a Core Graphics image object.
//
// https://developer.apple.com/library/ios/documentation/GraphicsImaging/Reference/CGImage/#//apple_ref/c/tdef/CGImageRef
type ImageRef C.CGImageRef

// ImageCreate creates a new Core Graphics image object that represents the
// same content than the Go image passed as argument.
//
// The image content is copied by the funciton, it's the program's responsibility
// to free the resources allocated by the returned ImageRef with a call to
// CFRelease.
//
// The function supports any image types defined in the standard image package,
// but will panic if the program attempts to create a ImageRef from an
// unsupported value.
//
// https://developer.apple.com/library/mac/documentation/GraphicsImaging/Reference/CGImage/
func ImageCreate(img image.Image) ImageRef {
	data := extractImageData(img)

	memory := C.CFDataCreate(
		nil,
		(*C.UInt8)(unsafe.Pointer(&data.pixels[0])),
		C.CFIndex(len(data.pixels)),
	)

	provider := C.CGDataProviderCreateWithCFData(memory)

	cgimg := C.CGImageCreate(
		C.size_t(data.width),
		C.size_t(data.height),
		C.size_t(data.bpc),
		C.size_t(data.bpp),
		C.size_t(data.stride),
		data.colors,
		data.info,
		provider,
		nil,
		false,
		C.kCGRenderingIntentDefault,
	)

	C.CFRelease(provider)
	C.CFRelease(memory)
	C.CFRelease(data.colors)
	return ImageRef(cgimg)
}

// ImageCreateNoCopy creates a new Core Graphics image object that represents the
// same content than the Go image passed as argument.
//
// The image content is shared between the Go and Core Graphics images, so the
// program must ensure that the image.Image value it passed to the function is
// referenced and unmodified for as long as the returned ImageRef is in use.
// It's the program's responsibility to free the resources allocated by the
// returned ImageRef with a call to CFRelease.
//
// The function supports any image types defined in the standard image package,
// but will panic if the program attempts to create a ImageRef from an
// unsupported value.
//
// https://developer.apple.com/library/mac/documentation/GraphicsImaging/Reference/CGImage/
func ImageCreateNoCopy(img image.Image) ImageRef {
	data := extractImageData(img)

	provider := C.CGDataProviderCreateWithData(
		nil,
		unsafe.Pointer(&data.pixels[0]),
		C.size_t(len(data.pixels)),
		nil,
	)

	cgimg := C.CGImageCreate(
		C.size_t(data.width),
		C.size_t(data.height),
		C.size_t(data.bpc),
		C.size_t(data.bpp),
		C.size_t(data.stride),
		data.colors,
		data.info,
		provider,
		nil,
		false,
		C.kCGRenderingIntentDefault,
	)

	C.CFRelease(provider)
	C.CFRelease(data.colors)
	return ImageRef(cgimg)
}

type imageData struct {
	pixels []byte
	bpc    int
	bpp    int
	stride int
	width  int
	height int
	colors C.CGColorSpaceRef
	info   C.CGBitmapInfo
}

func extractImageData(img image.Image) imageData {
	bounds := img.Bounds()

	switch i := img.(type) {
	case *image.RGBA:
		return imageData{
			pixels: i.Pix,
			bpc:    8,
			bpp:    32,
			stride: i.Stride,
			width:  bounds.Dx(),
			height: bounds.Dy(),
			colors: C.CGColorSpaceCreateDeviceRGB(),
			info:   C.CGBitmapInfo(C.kCGBitmapByteOrder32Big) | C.CGBitmapInfo(C.kCGImageAlphaPremultipliedLast),
		}

	case *image.NRGBA:
		return imageData{
			pixels: i.Pix,
			bpc:    8,
			bpp:    32,
			stride: i.Stride,
			width:  bounds.Dx(),
			height: bounds.Dy(),
			colors: C.CGColorSpaceCreateDeviceRGB(),
			info:   C.CGBitmapInfo(C.kCGBitmapByteOrder32Big) | C.CGBitmapInfo(C.kCGImageAlphaLast),
		}

	case *image.RGBA64:
		return imageData{
			pixels: i.Pix,
			bpc:    16,
			bpp:    64,
			stride: i.Stride,
			width:  bounds.Dx(),
			height: bounds.Dy(),
			colors: C.CGColorSpaceCreateDeviceRGB(),
			info:   C.CGBitmapInfo(C.kCGBitmapByteOrder32Big) | C.CGBitmapInfo(C.kCGImageAlphaPremultipliedLast),
		}

	case *image.NRGBA64:
		return imageData{
			pixels: i.Pix,
			bpc:    16,
			bpp:    64,
			stride: i.Stride,
			width:  bounds.Dx(),
			height: bounds.Dy(),
			colors: C.CGColorSpaceCreateDeviceRGB(),
			info:   C.CGBitmapInfo(C.kCGBitmapByteOrder32Big) | C.CGBitmapInfo(C.kCGImageAlphaLast),
		}

	case *image.Alpha:
		return imageData{
			pixels: i.Pix,
			bpc:    8,
			bpp:    8,
			stride: i.Stride,
			width:  bounds.Dx(),
			height: bounds.Dy(),
			colors: C.CGColorSpaceCreateDeviceGray(),
			info:   C.CGBitmapInfo(C.kCGImageAlphaNone),
		}

	case *image.Alpha16:
		return imageData{
			pixels: i.Pix,
			bpc:    16,
			bpp:    16,
			stride: i.Stride,
			width:  bounds.Dx(),
			height: bounds.Dy(),
			colors: C.CGColorSpaceCreateDeviceGray(),
			info:   C.CGBitmapInfo(C.kCGImageAlphaNone),
		}

	case *image.Gray:
		return imageData{
			pixels: i.Pix,
			bpc:    8,
			bpp:    8,
			stride: i.Stride,
			width:  bounds.Dx(),
			height: bounds.Dy(),
			colors: C.CGColorSpaceCreateDeviceGray(),
			info:   C.CGBitmapInfo(C.kCGImageAlphaNone),
		}

	case *image.Gray16:
		return imageData{
			pixels: i.Pix,
			bpc:    16,
			bpp:    16,
			stride: i.Stride,
			width:  bounds.Dx(),
			height: bounds.Dy(),
			colors: C.CGColorSpaceCreateDeviceGray(),
			info:   C.CGBitmapInfo(C.kCGImageAlphaNone),
		}

	case *image.CMYK:
		return imageData{
			pixels: i.Pix,
			bpc:    8,
			bpp:    32,
			stride: i.Stride,
			width:  bounds.Dx(),
			height: bounds.Dy(),
			colors: C.CGColorSpaceCreateDeviceCMYK(),
			info:   C.CGBitmapInfo(C.kCGBitmapByteOrder32Big) | C.CGBitmapInfo(C.kCGImageAlphaNone),
		}

	default:
		panic(fmt.Sprintf("%T: unsupported image format", img))
	}
}
