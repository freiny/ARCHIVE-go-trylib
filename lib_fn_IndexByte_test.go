package fStrings

import (
	"strings"
	"testing"
)

// func IndexByte(s string, c byte) int {
func IndexByte(s string, c uint8) int {
	return Index(s, string(c))
}

var indexByteTests = []struct {
	s   string
	c   byte
	out int
}{
	{"", byte(0), -1},
	{"a", byte(0), -1},
	{"a", byte('a'), 0},
	{"abc", byte('b'), 1},
	{"abcdecdef", byte('b'), 1},
	{"abcdecdef", byte('d'), 3},
	{"abcdecdef", byte('f'), 8},
}

func TestLibIndexByte(t *testing.T) {
	for _, test := range indexByteTests {
		actual := IndexByte(test.s, test.c)
		libout := strings.IndexByte(test.s, test.c)
		if actual != libout {
			t.Errorf("IndexByte(%q, %q) = %d ; %d wanted", test.s, test.c, actual, libout)
		}
	}
}

func TestIndexByte(t *testing.T) {
	for _, test := range indexByteTests {
		actual := IndexByte(test.s, test.c)
		if actual != test.out {
			t.Errorf("IndexByte(%q, %q) = %d ; %d wanted", test.s, test.c, actual, test.out)
		}
	}
}

func BenchmarkIndexByte(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IndexByte("abcdefghijklmnopqrstuvwxyz", byte('v'))
	}
}
