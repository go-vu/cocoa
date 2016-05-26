// +build darwin

package CT

// #cgo CFLAGS: -Wno-unused-parameter
// #cgo LDFLAGS: -framework CoreFoundation -framework CoreGraphics -framework CoreText
//
// #include <CoreText/CoreText.h>
// #include "font.h"
import "C"
import (
	"image"
	"unsafe"

	"github.com/go-vu/cocoa/CF"
	"github.com/go-vu/cocoa/CG"
)

// FontSymbolicTraits is an enumeration representing the style attributes of a font.
//
// https://developer.apple.com/library/mac/documentation/Carbon/Reference/CTFontDescriptorRef/#//apple_ref/c/tdef/CTFontSymbolicTraits
type FontSymbolicTraits int

// These constants are all the possible values of the FontSymbolicTraits
// enumeration.
const (
	FontItalicTrait      = FontSymbolicTraits(C.kCTFontItalicTrait)
	FontBoldTrait        = FontSymbolicTraits(C.kCTFontItalicTrait)
	FontExpandedTrait    = FontSymbolicTraits(C.kCTFontItalicTrait)
	FontCondensedTrait   = FontSymbolicTraits(C.kCTFontCondensedTrait)
	FontMonoSpaceTrait   = FontSymbolicTraits(C.kCTFontCondensedTrait)
	FontVerticalTrait    = FontSymbolicTraits(C.kCTFontVerticalTrait)
	FontUIOptimizedTrait = FontSymbolicTraits(C.kCTFontUIOptimizedTrait)
	FontClassMaskTrait   = FontSymbolicTraits(C.kCTFontClassMaskTrait)
)

// The FontRef type is an untyped reference to a Core Text font object.
//
// https://developer.apple.com/library/mac/documentation/Carbon/Reference/CTFontRef/#//apple_ref/c/tdef/CTFontRef
type FontRef unsafe.Pointer

// FontCreateWithName creates a new font object from a name, size and optional
// affine transformation.
//
// https://developer.apple.com/library/mac/documentation/Carbon/Reference/CTFontRef/#//apple_ref/c/func/CTFontCreateWithName
func FontCreateWithName(name CF.StringRef, size CG.Float, transform *CG.AffineTransform) FontRef {
	return FontRef(C.CTFontCreateWithName(
		C.CFStringRef(name),
		C.CGFloat(size),
		makeCGAffineTransform(transform),
	))
}

// FontCreateCopyWithSymbolicTraits makes a copy of an existing font object
// but allows the program to change the affine transformation or the style
// attributes of the font on the copy.
//
// https://developer.apple.com/library/mac/documentation/Carbon/Reference/CTFontRef/#//apple_ref/c/func/CTFontCreateCopyWithAttributes
func FontCreateCopyWithSymbolicTraits(font FontRef, size CG.Float, transform *CG.AffineTransform, traits FontSymbolicTraits, mask FontSymbolicTraits) FontRef {
	return FontRef(C.CTFontCreateCopyWithSymbolicTraits(
		C.CTFontRef(font),
		C.CGFloat(size),
		makeCGAffineTransform(transform),
		C.CTFontSymbolicTraits(traits),
		C.CTFontSymbolicTraits(mask),
	))
}

// FontCopyPostScriptName returns a copy of the font's post-script name.
//
// https://developer.apple.com/library/mac/documentation/Carbon/Reference/CTFontRef/#//apple_ref/c/func/CTFontCopyPostScriptName
func FontCopyPostScriptName(font FontRef) CF.StringRef {
	return CF.StringRef(C.CTFontCopyPostScriptName(C.CTFontRef(font)))
}

// FontCopyFamilyName returns a copy of the font's family name.
//
// https://developer.apple.com/library/mac/documentation/Carbon/Reference/CTFontRef/#//apple_ref/c/func/CTFontCopyFamilyName
func FontCopyFamilyName(font FontRef) CF.StringRef {
	return CF.StringRef(C.CTFontCopyFamilyName(C.CTFontRef(font)))
}

// FontCopyDisplayName returns a copy of the font's display name.
//
// https://developer.apple.com/library/mac/documentation/Carbon/Reference/CTFontRef/#//apple_ref/c/func/CTFontCopyDisplayName
func FontCopyDisplayName(font FontRef) CF.StringRef {
	return CF.StringRef(C.CTFontCopyDisplayName(C.CTFontRef(font)))
}

// FontCopyFullName returns a copy of the font's full name.
//
// https://developer.apple.com/library/mac/documentation/Carbon/Reference/CTFontRef/#//apple_ref/c/func/CTFontCopyFullName
func FontCopyFullName(font FontRef) CF.StringRef {
	return CF.StringRef(C.CTFontCopyFullName(C.CTFontRef(font)))
}

// FontGetAscent returns the ascent value of the font passed as argument.
//
// https://developer.apple.com/library/mac/documentation/Carbon/Reference/CTFontRef/#//apple_ref/c/func/CTFontGetAscent
func FontGetAscent(font FontRef) CG.Float {
	return CG.Float(C.CTFontGetAscent(C.CTFontRef(font)))
}

// FontGetDescent returns the descent value of the font passed as argument.
//
// https://developer.apple.com/library/mac/documentation/Carbon/Reference/CTFontRef/#//apple_ref/c/func/CTFontGetDescent
func FontGetDescent(font FontRef) CG.Float {
	return CG.Float(C.CTFontGetDescent(C.CTFontRef(font)))
}

// FontGetLeading returns the leading value of the font passed as argument.
//
// https://developer.apple.com/library/mac/documentation/Carbon/Reference/CTFontRef/#//apple_ref/c/func/CTFontGetLeading
func FontGetLeading(font FontRef) CG.Float {
	return CG.Float(C.CTFontGetLeading(C.CTFontRef(font)))
}

// FontGlyphDraw draws the font glyph representing the rune given as second
// argument into the alpha image at the specified position.
// The function returns true if the rune could be drawn, false otherwise, which
// measn the font had no representation of the rune.
func FontGlyphDraw(font FontRef, char rune, origin CG.Point, alpha *image.Alpha) bool {
	return bool(C.CTFontGlyphDraw__(
		C.CTFontRef(font),
		C.UTF32Char(char),
		makeCGPoint(origin),
		(*C.UInt8)(unsafe.Pointer(&alpha.Pix[0])),
		C.size_t(alpha.Stride),
		C.size_t(alpha.Rect.Dx()),
		C.size_t(alpha.Rect.Dy()),
	))
}

// FontGlyphAdvance returns the 'advance' of the glyph representing the rune
// given as second argument.
//
// https://developer.apple.com/library/mac/documentation/TextFonts/Conceptual/CocoaTextArchitecture/TypoFeatures/TextSystemFeatures.html
func FontGlyphAdvance(font FontRef, char rune) CG.Float {
	return CG.Float(C.CTFontGlyphAdvance__(C.CTFontRef(font), C.UTF32Char(char)))
}

// FontGlyphBounds returns the 'advance' and 'bounds' of the glyph
// representing the rune given as second argument.
//
// https://developer.apple.com/library/mac/documentation/TextFonts/Conceptual/CocoaTextArchitecture/TypoFeatures/TextSystemFeatures.html
func FontGlyphBounds(font FontRef, char rune) (advance CG.Float, bounds CG.Rect) {
	b := C.CGRect{}
	a := C.CTFontGlyphBounds__(C.CTFontRef(font), C.UTF32Char(char), &b)
	return CG.Float(a), makeRect(b)
}

// FontKern returns the ideal spacing to leave between the two characters
// passed as argument.
//
// https://developer.apple.com/library/mac/documentation/TextFonts/Conceptual/CocoaTextArchitecture/TypoFeatures/TextSystemFeatures.html
func FontKern(font FontRef, char0 rune, char1 rune) CG.Float {
	return CG.Float(C.CTFontKern__(C.CTFontRef(font), C.UTF32Char(char0), C.UTF32Char(char1)))
}

func makePoint(p C.CGPoint) CG.Point {
	return CG.Point{
		X: CG.Float(p.x),
		Y: CG.Float(p.y),
	}
}

func makeCGPoint(p CG.Point) C.CGPoint {
	return C.CGPoint{
		x: C.CGFloat(p.X),
		y: C.CGFloat(p.Y),
	}
}

func makeSize(s C.CGSize) CG.Size {
	return CG.Size{
		Width:  CG.Float(s.width),
		Height: CG.Float(s.height),
	}
}

func makeRect(r C.CGRect) CG.Rect {
	return CG.Rect{
		Origin: makePoint(r.origin),
		Size:   makeSize(r.size),
	}
}

func makeCGAffineTransform(t *CG.AffineTransform) *C.CGAffineTransform {
	if t == nil {
		return nil
	}
	return &C.CGAffineTransform{
		a:  C.CGFloat(t.A),
		b:  C.CGFloat(t.B),
		c:  C.CGFloat(t.C),
		d:  C.CGFloat(t.D),
		tx: C.CGFloat(t.Tx),
		ty: C.CGFloat(t.Ty),
	}
}
