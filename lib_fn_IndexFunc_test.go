package fStrings

import (
	"strings"
	"testing"
)

func IndexFunc(s string, f func(rune) bool) int {
	for i := 0; i < len(s); i++ {
		if f(rune(s[i])) {
			return i
		}
	}
	return -1
}

var indexFuncTests = []struct {
	s   string
	f   func(rune) bool
	out int
}{
	{"abc", func(r rune) bool {
		return r == rune('b')
	}, 1},
	{"abCdE", func(r rune) bool {
		s := string(r)
		return s == strings.ToUpper(s)
	}, 2},
	{"abcdE", func(r rune) bool {
		s := string(r)
		return s == strings.ToUpper(s)
	}, 4},
	{"abcde", func(r rune) bool {
		s := string(r)
		return s == strings.ToUpper(s)
	}, -1},
}

func TestLibIndexFunc(t *testing.T) {
	for _, test := range indexFuncTests {
		actual := IndexFunc(test.s, test.f)
		libout := strings.IndexFunc(test.s, test.f)
		if actual != libout {
			t.Errorf("IndexFunc(%q, %T) = %d, %d wanted", test.s, test.f, actual, libout)
		}
	}
}

func TestIndexFunc(t *testing.T) {
	for _, test := range indexFuncTests {
		actual := IndexFunc(test.s, test.f)
		if actual != test.out {
			t.Errorf("IndexFunc(%q, %T) = %d, %d wanted", test.s, test.f, actual, test.out)
		}
	}

}

func BenchmarkIndexFunc(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IndexFunc("abcdefghijklmnopqrstUVWXYZ", func(r rune) bool {
			return r == rune('X')
		})
	}
}
