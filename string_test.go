// +build darwin

package cocoa

import "testing"

func TestNewString(t *testing.T) {
	tests := []string{
		"",
		"0123456789",
		"Hello World!",
		"你好",
	}

	for _, test := range tests {
		s1 := NewString(test)
		s2 := s1.String()

		if s2 != test {
			t.Error("invalid string:", s2, "!=", test)
		}
	}
}

func TestNewStringWrap(t *testing.T) {
	tests := []string{
		"",
		"0123456789",
		"Hello World!",
		"你好",
	}

	for _, test := range tests {
		s1 := CFStringCreate(test)
		s2 := NewStringWrap(s1)

		if s2.Ref() != s1 {
			t.Error("invalid string ref:", s1, "!=", s2)
		}
	}
}

func TestStringRelease(t *testing.T) {
	s := NewString("")
	s.release()

	if s.Ref() != nil {
		t.Error("releasing didn't set the internal string reference to nil")
	}
}
