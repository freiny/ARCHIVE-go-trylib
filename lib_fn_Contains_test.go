package fStrings

import (
	"strings"
	"testing"
)

func Contains(s, sub string) bool {
	if Index(s, sub) != -1 {
		return true
	}
	return false
}

var containsTests = []struct {
	s, sub string
	out    bool
}{
	{"", "", true},
	{"", "a", false},
	{"a", "", true},
	{"abcde", "", true},
	{"abcde", "a", true},
	{"abcde", "c", true},
	{"abcde", "e", true},
	{"abcde", "ab", true},
	{"abcde", "bc", true},
	{"abcde", "de", true},
	{"abcde", "B", false},
	{"abcde", "bC", false},
}

func TestLibContains(t *testing.T) {
	for _, test := range containsTests {
		actual := Contains(test.s, test.sub)
		libout := strings.Contains(test.s, test.sub)
		if actual != libout {
			t.Errorf("Contains(%q, %q) = %t ; %t wanted", test.s, test.sub, actual, libout)
		}
	}
}

func TestContains(t *testing.T) {
	for _, test := range containsTests {
		actual := Contains(test.s, test.sub)
		if actual != test.out {
			t.Errorf("Contains(%q, %q) = %t ; %t wanted", test.s, test.sub, actual, test.out)
		}
	}
}

func BenchmarkContains(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Contains("abcdefghijklmnopqrstuvwxyz", "vwx")
	}
}
