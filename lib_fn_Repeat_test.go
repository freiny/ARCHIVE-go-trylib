package fStrings

import (
	"strings"
	"testing"
)

func Repeat(s string, count int) string {
	ret := ""
	for i := 0; i < count; i++ {
		ret += s
	}
	return ret
}

var repeatTests = []struct {
	s     string
	count int
	out   string
}{
	{"", 0, ""},
	{"", 2, ""},
	{"a", 0, ""},
	{"a", 2, "aa"},
	{"ab", 3, "ababab"},
	// {"ab", -2, ""},
}

func TestLibRepeat(t *testing.T) {
	for _, test := range repeatTests {
		actual := Repeat(test.s, test.count)
		libout := strings.Repeat(test.s, test.count)
		if actual != libout {
			t.Errorf("Repeat(%q, %d) = %q ; %q wanted", test.s, test.count, actual, libout)
		}
	}
}

func TestRepeat(t *testing.T) {
	for _, test := range repeatTests {
		actual := Repeat(test.s, test.count)
		if actual != test.out {
			t.Errorf("Repeat(%q, %d) = %q ; %q wanted", test.s, test.count, actual, test.out)
		}
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("ace-", 10)
	}
}
