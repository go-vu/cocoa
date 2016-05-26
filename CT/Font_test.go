// +build darwin

package CT

import (
	"testing"

	"github.com/go-vu/cocoa/CF"
)

func TestFontCreateWithName(t *testing.T) {
	s := CF.StringCreate("Monaco")
	f := FontCreateWithName(s, 12.0, nil)

	if f == nil {
		t.Error("failed to create a font")
		return
	}

	CF.Release(CF.TypeRef(f))
}
