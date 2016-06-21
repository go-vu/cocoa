// +build darwin

package CG

// #cgo CFLAGS: -Wno-unused-parameter
// #cgo LDFLAGS: -framework CoreFoundation -framework CoreGraphics
//
// #include <CoreGraphics/CoreGraphics.h>
import "C"

// Point is a Go equivalent to the type of the same name provided by Core
// Graphics.
//
// https://developer.apple.com/library/mac/documentation/GraphicsImaging/Reference/CGGeometry/#//apple_ref/c/tdef/CGPoint
type Point struct {
	X Float
	Y Float
}

// Size is a Go equivalent to the type of the same name provided by Core
// Graphics.
//
// https://developer.apple.com/library/mac/documentation/GraphicsImaging/Reference/CGGeometry/#//apple_ref/c/tdef/CGSize
type Size struct {
	Width  Float
	Height Float
}

// Rect is a Go equivalent to the type of the same name provided by Core
// Graphics.
//
// https://developer.apple.com/library/mac/documentation/GraphicsImaging/Reference/CGGeometry/#//apple_ref/c/tdef/CGRect
type Rect struct {
	Origin Point
	Size   Size
}

// The CGAffineTransform struct is a Go equivalent to the type of the same name
// provided by Core Graphics.
//
// https://developer.apple.com/library/mac/documentation/GraphicsImaging/Reference/CGAffineTransform/index.html#//apple_ref/doc/c_ref/CGAffineTransform
type AffineTransform struct {
	A  Float
	B  Float
	C  Float
	D  Float
	Tx Float
	Ty Float
}

// AffineTransformIdentity represents the identity matrix.
//
// https://developer.apple.com/library/mac/documentation/GraphicsImaging/Reference/CGAffineTransform/index.html#//apple_ref/doc/constant_group/CGAffineTransformIdentity
var AffineTransformIdentity = AffineTransform{
	1, 0, 0, 1, 0, 0,
}
