// +build darwin

package CT

import (
	"testing"

	"github.com/go-vu/cocoa/CF"
)

func TestFontCreateWithName(t *testing.T) {
	s := CF.StringCreate("Monaco")
	f := FontCreateWithName(s, 12.0, nil)

	defer s.Release()
	defer f.Release()

	if f == 0 {
		t.Error("failed to create a font")
	}
}

func TestFontCreateCopyWithSymbolicTraits(t *testing.T) {
	s := CF.StringCreate("Monaco")
	f := FontCreateWithName(s, 12.0, nil)
	g := FontCreateCopyWithSymbolicTraits(f, 20.0, nil, FontItalicTrait, FontClassMaskTrait)

	defer s.Release()
	defer f.Release()
	defer g.Release()

	if g == 0 {
		t.Error("failed to create font with symbol traits")
	}
}

func TestFontCopyPostScriptName(t *testing.T) {
	s := CF.StringCreate("Monaco")
	f := FontCreateWithName(s, 12.0, nil)

	defer s.Release()
	defer f.Release()

	if name := f.CopyPostScriptName(); name == 0 || name.String() != "Monaco" {
		t.Errorf("invalid post-script name:", name)
	}
}
