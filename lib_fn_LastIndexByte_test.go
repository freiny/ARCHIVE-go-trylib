package fStrings

import (
	"strings"
	"testing"
)

func LastIndexByte(s string, c byte) int {
	return LastIndex(s, string(c))
}

var lastIndexByteTests = []struct {
	s   string
	c   byte
	out int
}{
	{"", byte(0), -1},
	{"", byte('a'), -1},
	{"a", byte(0), -1},
	{"a", byte('a'), 0},
	{"a", byte('b'), -1},
	{"abcde", byte('a'), 0},
	{"abcde", byte('c'), 2},
	{"abcde", byte('e'), 4},
	{"abcde", byte('z'), -1},
	{"aaabbbccc", byte('a'), 2},
	{"aaabbbccc", byte('b'), 5},
	{"aaabbbccc", byte('c'), 8},
	{"aaabbbccc", byte('D'), -1},
}

func TestLibLastIndexByte(t *testing.T) {
	for _, test := range lastIndexByteTests {
		actual := LastIndexByte(test.s, test.c)
		libout := strings.LastIndexByte(test.s, test.c)
		if actual != libout {
			t.Errorf("LastIndexByte(%q, %q) = %d ; %d wanted", test.s, test.c, actual, libout)
		}
	}
}

func TestLastIndexByte(t *testing.T) {
	for _, test := range lastIndexByteTests {
		actual := LastIndexByte(test.s, test.c)
		if actual != test.out {
			t.Errorf("LastIndexByte(%q, %q) = %d ; %d wanted", test.s, test.c, actual, test.out)
		}
	}
}

func BenchmarkLastIndexByte(b *testing.B) {
	for i := 0; i < b.N; i++ {
		LastIndexByte("abcdefghijklmnopqrstuvwxyz", byte('c'))
	}
}
