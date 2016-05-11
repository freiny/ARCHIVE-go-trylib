package fStrings

import (
	"strings"
	"testing"
	"unicode"
)

func ToUpperSpecial(_case unicode.SpecialCase, s string) string {
	return Map(func(r rune) rune { return _case.ToUpper(r) }, s)
}

var toUpperSpecialTests = []struct {
	_case unicode.SpecialCase
	s     string
	out   string
}{
	{unicode.TurkishCase, "", ""},
	{unicode.TurkishCase, "i", "İ"},
	{unicode.TurkishCase, "ii", "İİ"},
}

func TestLibToUpperSpecial(t *testing.T) {
	for _, test := range toUpperSpecialTests {
		actual := ToUpperSpecial(test._case, test.s)
		libout := strings.ToUpperSpecial(test._case, test.s)
		if actual != libout {
			t.Errorf("ToUpperSpecial(%T, %q) = %q ; %q wanted", test._case, test.s, actual, libout)
		}
	}
}

func TestToUpperSpecial(t *testing.T) {
	for _, test := range toUpperSpecialTests {
		actual := ToUpperSpecial(test._case, test.s)
		if actual != test.out {
			t.Errorf("ToUpperSpecial(%T, %q) = %q ; %q wanted", test._case, test.s, actual, test.out)
		}
	}
}

func BenchmarkToUpperSpecial(b *testing.B) {
}
