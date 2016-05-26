// +build darwin

package CF

import "testing"

func TestString(t *testing.T) {
	tests := []string{
		"",
		"0123456789",
		"Hello World!",
		"你好",
	}

	for _, test := range tests {
		s1 := StringCreate(test)
		Retain(TypeRef(s1))
		Release(TypeRef(s1))

		s2 := GoString(s1)
		Release(TypeRef(s1))

		if test != s2 {
			t.Errorf(test, "!=", s2)
		}
	}
}
