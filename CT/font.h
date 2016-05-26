#ifndef GOVU_COCOA_FONT_H
#define GOVU_COCOA_FONT_H

#include <CoreGraphics/CoreGraphics.h>
#include <CoreText/CoreText.h>

bool CTFontGlyphDraw__(CTFontRef font, UTF32Char character, CGPoint origin,
                       UInt8 *buffer, size_t stride, size_t width,
                       size_t height);

CGFloat CTFontGlyphAdvance__(CTFontRef font, UTF32Char character);

CGFloat CTFontGlyphBounds__(CTFontRef font, UTF32Char character,
                            CGRect *bounds);

CGFloat CTFontKern__(CTFontRef font, UTF32Char from, UTF32Char to);

CGFloat CTFontKerningValueToPoints__(CTFontRef font, KernKerningValue kern);

#endif /* GOVU_COCOA_FONT_H */
