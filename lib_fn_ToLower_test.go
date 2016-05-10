package fStrings

import (
	"strings"
	"testing"
	"unicode"
)

func ToLower(s string) string {
	return Map(unicode.ToLower, s)
}

var toLowerTests = []struct {
	s   string
	out string
}{
	{"", ""},
	{" ", " "},
	{"\n", "\n"},
	{"a", "a"},
	{"A", "a"},
	{"b", "b"},
	{"B", "b"},
	{"3", "3"},
	{".", "."},
	{"\u00C0", "\u00E0"},
	{"\u00C1", "\u00E1"},
	{"\u00C2", "\u00E2"},
	{"\u00C3", "\u00E3"},
}

func TestLibToLower(t *testing.T) {
	for _, test := range toLowerTests {
		actual := ToLower(test.s)
		libout := strings.ToLower(test.s)
		if actual != libout {
			t.Errorf("ToLower(%q) = %q ; %q wanted", test.s, actual, libout)
		}
	}
}

func TestToLower(t *testing.T) {
	for _, test := range toLowerTests {
		actual := ToLower(test.s)
		if actual != test.out {
			t.Errorf("ToLower(%q) = %q ; %q wanted", test.s, actual, test.out)
		}
	}
}

func BenchmarkToLower(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ToLower("Hello World!")
	}
}
