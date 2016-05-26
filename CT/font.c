// +build darwin

#include "font.h"
#include "kern.h"

bool CTFontGlyphDraw__(CTFontRef font, UTF32Char character, CGPoint origin,
                       UInt8 *buffer, size_t stride, size_t width,
                       size_t height) {
  bool ok = false;

  CFStringRef string = CFStringCreateWithBytesNoCopy(
      NULL, (const UInt8 *)&character, sizeof(character),
      kCFStringEncodingUTF32LE, 0, kCFAllocatorNull);

  UniChar unichars[4] = {0};
  CGGlyph glyphs[4] = {0};
  CFIndex length = CFStringGetLength(string);
  CFStringGetCharacters(string, CFRangeMake(0, length), unichars);

  if (CTFontGetGlyphsForCharacters(font, unichars, glyphs, length)) {
    // The Quartz space has its origin in the bottom left corner, here we flip
    // the coordinate system to draw make the origin the top-left corner.
    origin.y = -(origin.y - height);

    CGPoint positions[4] = {origin, origin, origin, origin};

    for (CFIndex i = 1; i < length; ++i) {
      CGFloat offset = 0.0;
      CTFontGetLigatureCaretPositions(font, glyphs[i], &offset, 1);
      positions[i].x = positions[i - 1].x + offset;
    }

    CGColorSpaceRef colors = CGColorSpaceCreateDeviceGray();
    CGContextRef gc = CGBitmapContextCreateWithData(
        buffer, width, height, 8, stride, colors, 0, NULL, NULL);

    CGContextSetAllowsFontSubpixelPositioning(gc, true);
    CGContextSetShouldSubpixelPositionFonts(gc, true);
    CGContextSetGrayFillColor(gc, 1.0, 1.0);
    CTFontDrawGlyphs(font, glyphs, positions, length, gc);

    CGContextRelease(gc);
    CGColorSpaceRelease(colors);
    ok = true;
  }

  CFRelease(string);
  return ok;
}

CGFloat CTFontGlyphAdvance__(CTFontRef font, UTF32Char character) {
  CFStringRef string = CFStringCreateWithBytesNoCopy(
      NULL, (const UInt8 *)&character, sizeof(character),
      kCFStringEncodingUTF32LE, 0, kCFAllocatorNull);

  UniChar unichars[4] = {0};
  CGGlyph glyphs[4] = {0};
  CGFloat advance = 0.0;
  CFIndex length = CFStringGetLength(string);
  CFStringGetCharacters(string, CFRangeMake(0, length), unichars);

  if (CTFontGetGlyphsForCharacters(font, unichars, glyphs, length)) {
    advance = CTFontGetAdvancesForGlyphs(font, 0, glyphs, NULL, length);
  }

  CFRelease(string);
  return advance;
}

CGFloat CTFontGlyphBounds__(CTFontRef font, UTF32Char character,
                            CGRect *bounds) {
  CFStringRef string = CFStringCreateWithBytesNoCopy(
      NULL, (const UInt8 *)&character, sizeof(character),
      kCFStringEncodingUTF32LE, 0, kCFAllocatorNull);

  UniChar unichars[4] = {0};
  CGGlyph glyphs[4] = {0};
  CGFloat advance = 0.0;
  CFIndex length = CFStringGetLength(string);
  CFStringGetCharacters(string, CFRangeMake(0, length), unichars);

  if (CTFontGetGlyphsForCharacters(font, unichars, glyphs, length)) {
    CGRect rectangle =
        CTFontGetBoundingRectsForGlyphs(font, 0, glyphs, NULL, length);

    if (!CGRectIsNull(rectangle)) {
      *bounds = rectangle;
      advance = CTFontGetAdvancesForGlyphs(font, 0, glyphs, NULL, length);
    }
  }

  CFRelease(string);
  return advance;
}

CGFloat CTFontKern__(CTFontRef font, UTF32Char char0, UTF32Char char1) {
  CFStringRef string0 = CFStringCreateWithBytesNoCopy(
      NULL, (const UInt8 *)&char0, sizeof(char0), kCFStringEncodingUTF32LE, 0,
      kCFAllocatorNull);

  CFStringRef string1 = CFStringCreateWithBytesNoCopy(
      NULL, (const UInt8 *)&char1, sizeof(char1), kCFStringEncodingUTF32LE, 0,
      kCFAllocatorNull);

  UniChar unichars0[4] = {0};
  UniChar unichars1[4] = {0};
  CFIndex length0 = CFStringGetLength(string0);
  CFIndex length1 = CFStringGetLength(string1);
  CFStringGetCharacters(string0, CFRangeMake(0, length0), unichars0);
  CFStringGetCharacters(string1, CFRangeMake(0, length1), unichars1);

  CFDataRef table = CTFontCopyTable(font, 'kern', kCTFontTableOptionNoOptions);
  CGFloat kern = 0.0;

  if (table != NULL) {
    kern = CTFontKerningValueToPoints__(
        font,
        KernGet(CFDataGetBytePtr(table), unichars0[length0 - 1], unichars1[0]));
    CFRelease(table);
  }

  CFRelease(string1);
  CFRelease(string0);
  return kern;
}

CGFloat CTFontKerningValueToPoints__(CTFontRef font, KernKerningValue kern) {
  const CGAffineTransform tm = CTFontGetMatrix(font);
  const CGFloat unit = CTFontGetUnitsPerEm(font);
  const CGFloat size = CTFontGetSize(font);
  return (kern * size * tm.a) / unit;
}
