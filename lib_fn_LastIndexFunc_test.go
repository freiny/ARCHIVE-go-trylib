package fStrings

import (
	"strings"
	"testing"
)

func LastIndexFunc(s string, f func(rune) bool) int {
	for i := len(s) - 1; i >= 0; i-- {
		if f(rune(s[i])) {
			return i
		}
	}
	return -1
}

var lastIndexFuncTests = []struct {
	s   string
	f   func(rune) bool
	out int
}{
	// {"", nil, -1},
	// {"abc", nil, -1},

	{"", func(r rune) bool {
		return r == rune('a')
	}, -1},

	{"a", func(r rune) bool {
		return r == rune('a')
	}, 0},

	{"b", func(r rune) bool {
		return r == rune('a')
	}, -1},

	{"abc", func(r rune) bool {
		return r == rune('a')
	}, 0},

	{"def", func(r rune) bool {
		return r == rune('e')
	}, 1},

	{"ghi", func(r rune) bool {
		return r == rune('i')
	}, 2},

	{"aaabbbccc", func(r rune) bool {
		return r == rune('a')
	}, 2},

	{"dddeeefff", func(r rune) bool {
		return r == rune('e')
	}, 5},

	{"ggghhhiii", func(r rune) bool {
		return r == rune('i')
	}, 8},

	{"AbCdE", func(r rune) bool {
		s := string(r)
		return s == strings.ToUpper(s)
	}, 4},

	{"abcde", func(r rune) bool {
		s := string(r)
		return s == strings.ToUpper(s)
	}, -1},
}

func TestLibLastIndexFunc(t *testing.T) {
	for _, test := range lastIndexFuncTests {
		actual := LastIndexFunc(test.s, test.f)
		libout := strings.LastIndexFunc(test.s, test.f)
		if actual != libout {
			t.Errorf("LastIndexFunc(%q, %T) = %d ; %d wanted", test.s, test.f, actual, libout)
		}
	}
}

func TestLastIndexFunc(t *testing.T) {
	for _, test := range lastIndexFuncTests {
		actual := LastIndexFunc(test.s, test.f)
		if actual != test.out {
			t.Errorf("LastIndexFunc(%q, %T) = %d ; %d wanted", test.s, test.f, actual, test.out)
		}
	}
}

func BenchmarkLastIndexFunc(b *testing.B) {
	for i := 0; i < b.N; i++ {
		LastIndexFunc("abcdefghijkLmnopqrstuvwxyz", func(r rune) bool {
			s := string(r)
			return s == strings.ToUpper(s)
		})
	}
}
