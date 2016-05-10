package fStrings

import (
	"strings"
	"testing"
)

func Compare(a, b string) int {
	if a > b {
		return 1
	}
	if a < b {
		return -1
	}
	return 0
}

var compareTests = []struct {
	a, b string
	out  int
}{
	{"", "", 0},
	{"a", "", 1},
	{"", "a", -1},
	{"a", "a", 0},
	{"a", "b", -1},
	{"b", "a", 1},
	{"ab", "ab", 0},
	{"ab", "bc", -1},
	{"bc", "ab", 1},
	{"ab", "abc", -1},
	{"abc", "bc", -1},
}

func TestLibCompare(t *testing.T) {

	for _, test := range compareTests {
		actual := Compare(test.a, test.b)
		libout := strings.Compare(test.a, test.b)
		if actual != libout {
			t.Errorf("Compare(%q, %q) = %d ; %d wanted", test.a, test.b, actual, libout)
		}
	}

}

func TestCompare(t *testing.T) {

	for _, test := range compareTests {
		actual := Compare(test.a, test.b)
		if actual != test.out {
			t.Errorf("Compare(%q, %q) = %d ; %d wanted", test.a, test.b, actual, test.out)
		}
	}

}

func BenchmarkCompare(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Compare("abc", "bc")
	}
}
