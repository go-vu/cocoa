// +build darwin

package cocoa

import "testing"

func TestCFString(t *testing.T) {
	tests := []string{
		"",
		"0123456789",
		"Hello World!",
		"你好",
	}

	for _, test := range tests {
		s1 := CFStringCreate(test)
		CFStringRetain(s1)
		CFStringRelease(s1)

		s2 := GoString(s1)
		CFStringRelease(s1)

		if test != s2 {
			t.Errorf(test, "!=", s2)
		}
	}
}
