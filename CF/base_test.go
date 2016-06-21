package CF

import "testing"

func TestTypeRefString(t *testing.T) {
	s1 := StringCreate("Hello World!")
	s2 := TypeRef(s1).String()

	if len(s2) == 0 {
		t.Error("invalid string description:", s2)
	}

	TypeRef(s1).Release()
}

func TestTypeIDString(t *testing.T) {
	s1 := StringCreate("Hello World!")
	s2 := TypeRef(s1).GetTypeID().String()

	if s2 != "CFString" {
		t.Error("invalid string description:", s2)
	}

	TypeRef(s1).Release()

}
