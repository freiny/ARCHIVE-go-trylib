package fStrings

import (
	"strings"
	"testing"
)

func LastIndexAny(s, chars string) int {
	switch {
	case s == "" || chars == "":
		return -1
	default:

		currMax := -1
		for i := 0; i < len(chars); i++ {
			li := LastIndex(s, string(chars[i]))
			if li > currMax {
				currMax = li
			}
		}
		return currMax
	}
}

var lastIndexAnyTests = []struct {
	s, chars string
	out      int
}{
	{"", "", -1},
	{"", "a", -1},
	{"a", "", -1},
	{"a", "a", 0},
	{"abcde", "", -1},
	{"abcde", "abcde", 4},
	{"abcde", "edcba", 4},
	{"aaabbbccc", "bc", 8},
	{"aaabbbccc", "cb", 8},
	{"aaabbbccc", "aa", 2},
	{"aaabbbccc", "aba", 5},
	{"aaabbbccc", "b", 5},
	{"aaabbbccc", "c", 8},
	{"aaabbbccc", "eiou", -1},
}

func TestLibLastIndexAny(t *testing.T) {
	for _, test := range lastIndexAnyTests {
		actual := LastIndexAny(test.s, test.chars)
		libout := strings.LastIndexAny(test.s, test.chars)
		if actual != libout {
			t.Errorf("LastIndexAny(%q, %q) = %d ; %d wanted", test.s, test.chars, actual, libout)
		}
	}
}

func TestLastIndexAny(t *testing.T) {
	for _, test := range lastIndexAnyTests {
		actual := LastIndexAny(test.s, test.chars)
		if actual != test.out {
			t.Errorf("LastIndexAny(%q, %q) = %d ; %d wanted", test.s, test.chars, actual, test.out)
		}
	}
}

func BenchmarkLastIndexAny(b *testing.B) {
	for i := 0; i < b.N; i++ {
		LastIndexAny("aaabbbcccdddeee", "abcde")
	}
}
