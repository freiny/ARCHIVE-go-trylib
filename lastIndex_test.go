package fStrings

import (
	"strings"
	"testing"
)

func LastIndex(s, sep string) int {

	sp := s[0:len(s)]

	switch {
	case s == sep:
		return 0
	case len(s) > len(sep):
		if sep == "" {
			return len(s)
		}
		for i := len(s) - len(sep); i > 0; i-- {
			if sp[i:i+len(sep)] == sep {
				return i
			}
		}
	}

	return -1
}

var lastIndexTests = []struct {
	s, sep string
	out    int
}{
	{"", "", 0},
	{"", "a", -1},
	{"a", "", 1},
	{"a", "a", 0},
	{"abcde", "", 5},
	{"abcde", "abcde", 0},
	{"aaabbbccc", "a", 2},
	{"aaabbbccc", "b", 5},
	{"aaabbbccc", "c", 8},
	{"aaabbbccc", "aa", 1},
	{"aaabbbccc", "bb", 4},
	{"aaabbbccc", "cc", 7},
	{"aaa aa aaa a aaa", "aaa", 13},
	{"aaabbbccc", "k", -1},
}

func TestLibLastIndex(t *testing.T) {
	for _, test := range lastIndexTests {
		actual := LastIndex(test.s, test.sep)
		libOut := strings.LastIndex(test.s, test.sep)
		if actual != libOut {
			t.Errorf("LastIndex(%q, %q) = %d ; %d", test.s, test.sep, actual, libOut)
		}
	}
}

func TestLastIndex(t *testing.T) {
	for _, test := range lastIndexTests {
		actual := LastIndex(test.s, test.sep)
		if actual != test.out {
			t.Errorf("LastIndex(%q, %q) = %d ; %d", test.s, test.sep, actual, test.out)
		}
	}
}

func BenchmarkLastIndex(b *testing.B) {
	for i := 0; i < b.N; i++ {
		LastIndex("aaaabbbbccccddddeeee", "ee")
	}
}
