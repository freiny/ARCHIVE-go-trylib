package fStrings

import (
	"strings"
	"testing"
)

// TODO: implement EqualFold

func EqualFold(s, t string) bool {
	return strings.EqualFold(s, t)
}

var equalFoldTests = []struct {
	s, t string
	out  bool
}{
	{"aBcDeFg", "abcdefg", true},
	{"aBcDeFg", "abc", false},
}

func TestLibEqualFold(t *testing.T) {
	for _, test := range equalFoldTests {
		actual := EqualFold(test.s, test.t)
		libout := strings.EqualFold(test.s, test.t)
		if actual != libout {
			t.Errorf("EqualFold(%q, %q) = %t ; %t", test.s, test.t, actual, libout)
		}
	}
}

func TestEqualFold(t *testing.T) {
	for _, test := range equalFoldTests {
		actual := EqualFold(test.s, test.t)
		if actual != test.out {
			t.Errorf("EqualFold(%q, %q) = %t ; %t", test.s, test.t, actual, test.out)
		}
	}
}

func BenchmarkEqualFold(b *testing.B) {
	for i := 0; i < b.N; i++ {
		EqualFold("aBcDeFg", "abcdefg")
	}
}
