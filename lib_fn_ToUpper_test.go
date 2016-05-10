package fStrings

import (
	"strings"
	"testing"
	"unicode"
)

func ToUpper(s string) string {
	return Map(unicode.ToUpper, s)
}

var toUpperTests = []struct {
	s   string
	out string
}{
	{"", ""},
	{" ", " "},
	{"\n", "\n"},
	{"a", "A"},
	{"A", "A"},
	{"b", "B"},
	{"B", "B"},
	{"3", "3"},
	{".", "."},
	{"\u00E0", "\u00C0"},
	{"\u00E1", "\u00C1"},
	{"\u00E2", "\u00C2"},
	{"\u00E3", "\u00C3"},
}

func TestLibToUpper(t *testing.T) {
	for _, test := range toUpperTests {
		actual := ToUpper(test.s)
		libout := strings.ToUpper(test.s)
		if actual != libout {
			t.Errorf("ToUpper(%q) = %q ; %q wanted", test.s, actual, libout)
		}
	}
}

func TestToUpper(t *testing.T) {
	for _, test := range toUpperTests {
		actual := ToUpper(test.s)
		if actual != test.out {
			t.Errorf("ToUpper(%q) = %q ; %q wanted", test.s, actual, test.out)
		}
	}
}

func BenchmarkToUpper(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ToUpper("Hello World!")
	}
}
