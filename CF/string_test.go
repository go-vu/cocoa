// +build darwin

package CF

import (
	"fmt"
	"testing"
)

var strings = []string{
	"",
	"0123456789",
	"Hello World!",
	"你好",
	"\"abc\"\n",
}

func TestString(t *testing.T) {
	for _, test := range strings {
		s1 := StringCreate(test)
		s1.Retain()
		s1.Release()

		s2 := GoString(s1)
		s1.Release()

		if test != s2 {
			t.Error(test, "!=", s2)
		}
	}
}

func TestStringEqualTrue(t *testing.T) {
	s1 := StringCreate("Hello World!")
	s2 := StringCreate("Hello World!")
	defer s1.Release()
	defer s2.Release()

	if !Equal(TypeRef(s1), TypeRef(s2)) {
		t.Error("comparing strings for equality failed")
	}
}

func TestStringEqualFalse(t *testing.T) {
	s1 := StringCreate("Hello World!")
	s2 := StringCreate("A")
	defer s1.Release()
	defer s2.Release()

	if Equal(TypeRef(s1), TypeRef(s2)) {
		t.Error("comparing strings for difference failed")
	}
}

func TestStringString(t *testing.T) {
	for _, test := range strings {
		s1 := StringCreate(test)
		s2 := s1.String()
		s1.Release()

		if test != s2 {
			t.Error(test, "!=", s2)
		}
	}
}

func TestStringGoString(t *testing.T) {
	for _, test := range strings {
		s1 := StringCreate(test)
		s2 := s1.GoString()
		s1.Release()

		if test != fmt.Sprintf("%v", test) {
			t.Error(test, "!=", s2)
		}
	}
}

func TestStringTypeID(t *testing.T) {
	s := StringCreate("Hello World!")
	id := TypeRef(s).GetTypeID()

	if id == 0 {
		t.Error("invalid zero type id got from string")
	}

	TypeRef(s).Release()
}

func TestStringCopyDescription(t *testing.T) {
	s1 := StringCreate("Hello World!")
	s2 := TypeRef(s1).CopyDescription()

	if s2.Length() == 0 {
		t.Error("invalid string description:", s2)
	}

	s1.Release()
	s2.Release()
}

func TestStringCopyTypeIDDescription(t *testing.T) {
	s1 := StringCreate("Hello World!")
	s2 := TypeRef(s1).GetTypeID().CopyTypeIDDescription()

	if !Equal(TypeRef(s2), TypeRef(StringCreate("CFString"))) {
		t.Error("invalid string description:", s2)
	}

	s1.Release()
	s2.Release()
}
