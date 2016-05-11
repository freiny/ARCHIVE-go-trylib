package fStrings

import (
	"strings"
	"testing"
	"unicode"
)

func ToLowerSpecial(_case unicode.SpecialCase, s string) string {
	return Map(func(r rune) rune { return _case.ToLower(r) }, s)
}

var toLowerSpecialTests = []struct {
	_case unicode.SpecialCase
	s     string
	out   string
}{
	{unicode.TurkishCase, "", ""},
	{unicode.TurkishCase, "İ", "i"},
	{unicode.TurkishCase, "İİ", "ii"},
}

func TestLibToLowerSpecial(t *testing.T) {
	for _, test := range toLowerSpecialTests {
		actual := ToLowerSpecial(test._case, test.s)
		libout := strings.ToLowerSpecial(test._case, test.s)
		if actual != libout {
			t.Errorf("ToLowerSpecial(%T, %q) = %q ; %q wanted", test._case, test.s, actual, libout)
		}
	}
}

func TestToLowerSpecial(t *testing.T) {
	for _, test := range toLowerSpecialTests {
		actual := ToLowerSpecial(test._case, test.s)
		if actual != test.out {
			t.Errorf("ToLowerSpecial(%T, %q) = %q ; %q wanted", test._case, test.s, actual, test.out)
		}
	}
}

func BenchmarkToLowerSpecial(b *testing.B) {
}
