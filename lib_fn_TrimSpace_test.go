package fStrings

import (
	"strings"
	"testing"
)

func TrimSpace(s string) string {
	cutset := "\t\n\v\f\r \u0085\u00A0"
	t := TrimLeft(s, cutset)
	return TrimRight(t, cutset)
}

var trimSpaceTests = []struct {
	s   string
	out string
}{
	{"", ""},
	{"a", "a"},

	{"\t", ""},
	{"\n", ""},
	{"\v", ""},
	{"\f", ""},
	{"\r", ""},
	{" ", ""},
	{"\u0085", ""},
	{"\u00A0", ""},

	{"a\t", "a"},
	{"a\n", "a"},
	{"a\v", "a"},
	{"a\f", "a"},
	{"a\r", "a"},
	{"a ", "a"},
	{"a\u0085", "a"},
	{"a\u00A0", "a"},

	{"\tb", "b"},
	{"\nb", "b"},
	{"\vb", "b"},
	{"\fb", "b"},
	{"\rb", "b"},
	{" b", "b"},
	{"\u0085b", "b"},
	{"\u00A0b", "b"},

	{"a\tb", "a\tb"},
	{"a\nb", "a\nb"},
	{"a\vb", "a\vb"},
	{"a\fb", "a\fb"},
	{"a\rb", "a\rb"},
	{"a b", "a b"},
	{"a\u0085b", "a\u0085b"},
	{"a\u00A0b", "a\u00A0b"},

	{"\ta b", "a b"},
	{"a b\t", "a b"},
	{"\na b\t", "a b"},
	{"a\n \tb", "a\n \tb"},
}

func TestLibTrimSpace(t *testing.T) {
	for _, test := range trimSpaceTests {
		actual := TrimSpace(test.s)
		libout := strings.TrimSpace(test.s)
		if actual != libout {
			t.Errorf("TrimSpace(%q) = %q ; %q wanted", test.s, actual, libout)
		}
	}
}

func TestTrimSpace(t *testing.T) {
	for _, test := range trimSpaceTests {
		actual := TrimSpace(test.s)
		if actual != test.out {
			t.Errorf("TrimSpace(%q) = %q ; %q wanted", test.s, actual, test.out)
		}
	}
}

func BenchmarkTrimSpace(b *testing.B) {
	for i := 0; i < b.N; i++ {
		TrimSpace("\n \va aa bb b\t \t\n")
	}
}
