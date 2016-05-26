#ifndef GOVU_DRIVER_COCOA_KERN_H
#define GOVU_DRIVER_COCOA_KERN_H

#include <CoreText/SFNTLayoutTypes.h>

KernKerningValue KernGet(const UInt8 *table, UniChar char0, UniChar char1);

KernKerningValue Kern0Get(const KernVersion0Header *header, UniChar char0,
                          UniChar char1);

KernKerningValue Kern1Get(const KernTableHeader *header, UniChar char0,
                          UniChar char1);

KernKerningValue KernOrderedListGet(const KernOrderedListHeader *header,
                                    UniChar char0, UniChar char1);

UInt32 KernMakeKey(UniChar left, UniChar right);

#endif /* GOVU_DRIVER_COCOA_KERN_H */
