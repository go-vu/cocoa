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
	FontItalicTrait      FontSymbolicTraits = FontSymbolicTraits(C.kCTFontItalicTrait)
	FontBoldTrait        FontSymbolicTraits = FontSymbolicTraits(C.kCTFontItalicTrait)
	FontExpandedTrait    FontSymbolicTraits = FontSymbolicTraits(C.kCTFontItalicTrait)
	FontCondensedTrait   FontSymbolicTraits = FontSymbolicTraits(C.kCTFontCondensedTrait)
	FontMonoSpaceTrait   FontSymbolicTraits = FontSymbolicTraits(C.kCTFontCondensedTrait)
	FontVerticalTrait    FontSymbolicTraits = FontSymbolicTraits(C.kCTFontVerticalTrait)
	FontUIOptimizedTrait FontSymbolicTraits = FontSymbolicTraits(C.kCTFontUIOptimizedTrait)
	FontClassMaskTrait   FontSymbolicTraits = FontSymbolicTraits(C.kCTFontClassMaskTrait)
)

// The FontRef type is an untyped reference to a Core Text font object.
//
// https://developer.apple.com/library/mac/documentation/Carbon/Reference/CTFontRef/#//apple_ref/c/tdef/CTFontRef
type FontRef CF.TypeRef

// FontCreateWithName creates a new font object from a name, size and optional
// affine transformation.
//
// https://developer.apple.com/library/mac/documentation/Carbon/Reference/CTFontRef/#//apple_ref/c/func/CTFontCreateWithName
func FontCreateWithName(name CF.StringRef, size CG.Float, transform *CG.AffineTransform) FontRef {
	return FontRef(unsafe.Pointer(C.CTFontCreateWithName(
		C.CFStringRef(unsafe.Pointer(name)),
		C.CGFloat(size),
		makeCGAffineTransform(transform),
	)))
}

// FontCreateCopyWithSymbolicTraits makes a copy of an existing font object
// but allows the program to change the affine transformation or the style
// attributes of the font on the copy.
//
// https://developer.apple.com/library/mac/documentation/Carbon/Reference/CTFontRef/#//apple_ref/c/func/CTFontCreateCopyWithAttributes
func FontCreateCopyWithSymbolicTraits(font FontRef, size CG.Float, transform *CG.AffineTransform, traits FontSymbolicTraits, mask FontSymbolicTraits) FontRef {
	return FontRef(unsafe.Pointer(C.CTFontCreateCopyWithSymbolicTraits(
		C.CTFontRef(unsafe.Pointer(font)),
		C.CGFloat(size),
		makeCGAffineTransform(transform),
		C.CTFontSymbolicTraits(traits),
		C.CTFontSymbolicTraits(mask),
	)))
}

// FontCopyPostScriptName returns a copy of the font's post-script name.
//
// https://developer.apple.com/library/mac/documentation/Carbon/Reference/CTFontRef/#//apple_ref/c/func/CTFontCopyPostScriptName
func (f FontRef) CopyPostScriptName() CF.StringRef {
	return CF.StringRef(unsafe.Pointer(C.CTFontCopyPostScriptName(C.CTFontRef(unsafe.Pointer(f)))))
}

// FontCopyFamilyName returns a copy of the font's family name.
//
// https://developer.apple.com/library/mac/documentation/Carbon/Reference/CTFontRef/#//apple_ref/c/func/CTFontCopyFamilyName
func (f FontRef) CopyFamilyName() CF.StringRef {
	return CF.StringRef(unsafe.Pointer(C.CTFontCopyFamilyName(C.CTFontRef(unsafe.Pointer(f)))))
}

// FontCopyDisplayName returns a copy of the font's display name.
//
// https://developer.apple.com/library/mac/documentation/Carbon/Reference/CTFontRef/#//apple_ref/c/func/CTFontCopyDisplayName
func (f FontRef) CopyDisplayName() CF.StringRef {
	return CF.StringRef(unsafe.Pointer(C.CTFontCopyDisplayName(C.CTFontRef(unsafe.Pointer(f)))))
}

// FontCopyFullName returns a copy of the font's full name.
//
// https://developer.apple.com/library/mac/documentation/Carbon/Reference/CTFontRef/#//apple_ref/c/func/CTFontCopyFullName
func (f FontRef) CopyFullName() CF.StringRef {
	return CF.StringRef(unsafe.Pointer(C.CTFontCopyFullName(C.CTFontRef(unsafe.Pointer(f)))))
}

// FontGetAscent returns the ascent value of the font passed as argument.
//
// https://developer.apple.com/library/mac/documentation/Carbon/Reference/CTFontRef/#//apple_ref/c/func/CTFontGetAscent
func (f FontRef) GetAscent() CG.Float {
	return CG.Float(C.CTFontGetAscent(C.CTFontRef(unsafe.Pointer(f))))
}

// FontGetDescent returns the descent value of the font passed as argument.
//
// https://developer.apple.com/library/mac/documentation/Carbon/Reference/CTFontRef/#//apple_ref/c/func/CTFontGetDescent
func (f FontRef) GetDescent() CG.Float {
	return CG.Float(C.CTFontGetDescent(C.CTFontRef(unsafe.Pointer(f))))
}

// FontGetLeading returns the leading value of the font passed as argument.
//
// https://developer.apple.com/library/mac/documentation/Carbon/Reference/CTFontRef/#//apple_ref/c/func/CTFontGetLeading
func (f FontRef) GetLeading() CG.Float {
	return CG.Float(C.CTFontGetLeading(C.CTFontRef(unsafe.Pointer(f))))
}

// FontGlyphDraw draws the font glyph representing the rune given as second
// argument into the alpha image at the specified position.
// The function returns true if the rune could be drawn, false otherwise, which
// measn the font had no representation of the rune.
func (f FontRef) GlyphDraw(char rune, origin CG.Point, alpha *image.Alpha) bool {
	return bool(C.CTFontGlyphDraw__(
		C.CTFontRef(unsafe.Pointer(f)),
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
func (f FontRef) GlyphAdvance(char rune) CG.Float {
	return CG.Float(C.CTFontGlyphAdvance__(C.CTFontRef(unsafe.Pointer(f)), C.UTF32Char(char)))
}

// FontGlyphBounds returns the 'advance' and 'bounds' of the glyph
// representing the rune given as second argument.
//
// https://developer.apple.com/library/mac/documentation/TextFonts/Conceptual/CocoaTextArchitecture/TypoFeatures/TextSystemFeatures.html
func (f FontRef) GlyphBounds(char rune) (advance CG.Float, bounds CG.Rect) {
	b := C.CGRect{}
	a := C.CTFontGlyphBounds__(C.CTFontRef(unsafe.Pointer(f)), C.UTF32Char(char), &b)
	return CG.Float(a), makeRect(b)
}

// FontKern returns the ideal spacing to leave between the two characters
// passed as argument.
//
// https://developer.apple.com/library/mac/documentation/TextFonts/Conceptual/CocoaTextArchitecture/TypoFeatures/TextSystemFeatures.html
func (f FontRef) Kern(char0 rune, char1 rune) CG.Float {
	return CG.Float(C.CTFontKern__(C.CTFontRef(unsafe.Pointer(f)), C.UTF32Char(char0), C.UTF32Char(char1)))
}

// Retain increases the refence counter of the Core Text font passed
// as argument.
//
// https://developer.apple.com/library/mac/documentation/CoreFoundation/Reference/CFTypeRef/#//apple_ref/c/func/CFRetain
func (f FontRef) Retain() {
	CF.TypeRef(f).Retain()
}

// Release decreases the reference counter of the Core Text font
// passed as argument.
//
// https://developer.apple.com/library/mac/documentation/CoreFoundation/Reference/CFTypeRef/#//apple_ref/c/func/CFRelease
func (f FontRef) Release() {
	CF.TypeRef(f).Release()
}

// String satisfies the fmt.Stringer interface.
func (f FontRef) String() string {
	return CF.TypeRef(f).String()
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
