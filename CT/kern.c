// +build darwin

#include <stdio.h>
#include "kern.h"

enum {
  Kern0Horizontal  = 1 << 0,
  Kern0Minimum     = 1 << 1,
  Kern0CrossStream = 1 << 2,
  Kern0Override    = 1 << 3,
  Kern0FormatMask  = 0xFF00,
};

KernKerningValue KernGet(const UInt8 *table, UniChar char0, UniChar char1) {
  const UInt16 version = *(const UInt16 *)(table);
  KernKerningValue kern;

  if (version == 0) {
    kern = Kern0Get((const KernVersion0Header *) table, char0, char1);
  } else {
    kern = Kern1Get((const KernTableHeader *) table, char0, char1);
  }

  return kern;
}

KernKerningValue Kern0Get(const KernVersion0Header *header, UniChar char0, UniChar char1) {
  KernKerningValue kern = 0;
  KernKerningValue value = 0;

  const KernVersion0SubtableHeader *subtable = (const KernVersion0SubtableHeader *) (header->firstSubtable);

  for (int i = 0, n = ntohs(header->nTables); i != n; i++) {
    const UInt16 version = ntohs(subtable->version);
    const UInt16 coverage = ntohs(subtable->stInfo);
    const UInt16 length = ntohs(subtable->length);

    if (version != 0) {
      fprintf(stderr, "WARN: kerning is not supported for tables with version %u\n", version);
      goto next;
    }

    if (!(coverage & Kern0Horizontal)) {
      fprintf(stderr, "WARN: kerning is not supported for `vertical` tables\n");
      goto next;
    }

    if (coverage & Kern0CrossStream) {
      fprintf(stderr, "WARN: kerning is not supported for tables `cross-stream` tables\n");
      goto next;
    }

    if (coverage & Kern0Minimum) {
      fprintf(stderr, "WARN: kerning is not supported for `minimum` tables\n");
      goto next;
    }

    const UInt16 format = (coverage & Kern0FormatMask) >> 8;

    switch (format) {
    case kKERNOrderedList:
      value = KernOrderedListGet(&subtable->fsHeader.orderedList, char0, char1);
      break;

    default:
      fprintf(stderr, "WARN: kerning is not supported for tables with format %u\n", format);
      goto next;
    }

    if (coverage & Kern0Override) {
      kern = value;
    } else {
      kern += value;
    }

  next:
    subtable = (const KernVersion0SubtableHeader *) (((const UInt8 *) subtable) + (length - 6));
  }

  return kern;
}

KernKerningValue Kern1Get(const KernTableHeader *header, UniChar char0, UniChar char1) {
  KernKerningValue kern = 0;

  const KernSubtableHeader *subtable = (const KernSubtableHeader *) (header->firstSubtable);

  for (int i = 0, n = ntohs(header->nTables); i != n; i++) {
    const UInt16 length = ntohs(subtable->length);
    const UInt16 coverage = subtable->stInfo;

    if (coverage & kKERNVertical) {
      fprintf(stderr, "WARN: kerning is not supported for `vertical` tables\n");
      goto next;
    }

    if (coverage & kKERNCrossStream) {
      fprintf(stderr, "WARN: kerning is not supported for tables `cross-stream` tables\n");
      goto next;
    }

    if (coverage & kKERNVariation) {
      fprintf(stderr, "WARN: kerning is not supported for `variation` tables\n");
      goto next;
    }

    const UInt16 format = coverage & kKERNFormatMask;

    switch (format) {
    case kKERNOrderedList:
      kern += KernOrderedListGet(&subtable->fsHeader.orderedList, char0, char1);
      break;

    default:
      fprintf(stderr, "WARN: kerning is not supported for tables with format %u\n", format);
      goto next;
    }

  next:
    subtable = (const KernSubtableHeader *) (((const UInt8 *) subtable) + (length - 8));
  }

  return kern;
}

KernKerningValue KernOrderedListGet(const KernOrderedListHeader *header, UniChar char0, UniChar char1) {
  const KernOrderedListEntry *ptr = (const KernOrderedListEntry *) (header->table);
  const KernOrderedListEntry *end = ptr + ntohs(header->nPairs);
  const KernOrderedListEntry *mid = ptr + ((end - ptr) >> 1);
  const UInt32 key = KernMakeKey(char0, char1);

  while (ptr != end) {
    const UInt32 val = KernMakeKey(ntohs(mid->pair.left), ntohs(mid->pair.right));

    if (key == val) {
      return ntohs(mid->value);
    }

    if ((end - ptr) == 1) {
      break;
    }

    if (key < val) {
      end = mid;
    } else {
      ptr = mid;
    }

    mid = ptr + ((end - ptr) >> 1);
  }

  return 0;
}

UInt32 KernMakeKey(UniChar left, UniChar right) {
  const UInt32 a = left;
  const UInt32 b = right;
  return (a << 16) | b;
}
