package fStrings

import (
	"strings"
	"testing"
)

func IndexAny(s, chars string) int {
	ret := len(s)
	for i := 0; i < len(chars); i++ {
		// fmt.Println(s, string(chars[i]), Index(s, string(chars[i])))
		index := Index(s, string(chars[i]))
		if index < ret && index != -1 {
			ret = index
		}
	}

	if ret == len(s) {
		ret = -1
	}
	return ret
}

var indexAnyTests = []struct {
	s, chars string
	out      int
}{
	{"", "", -1},
	{"", "a", -1},
	{"a", "", -1},
	{"a", "a", 0},
	{"abcde", "abcde", 0},
	{"abcde", "c", 2},
	{"abcde", "dcb", 1},
	{"abcde", "bde", 1},
	{"abcde", "wxyzbcdq", 1},
}

func TestLibIndexAny(t *testing.T) {
	for _, test := range indexAnyTests {
		actual := IndexAny(test.s, test.chars)
		libout := strings.IndexAny(test.s, test.chars)
		if actual != libout {
			t.Errorf("IndexAny(%q, %q) = %d ; %d wanted", test.s, test.chars, actual, libout)
		}
	}
}

func TestIndexAny(t *testing.T) {
	for _, test := range indexAnyTests {
		actual := IndexAny(test.s, test.chars)
		if actual != test.out {
			t.Errorf("IndexAny(%q, %q) = %d ; %d wanted", test.s, test.chars, actual, test.out)
		}
	}
}

func BenchmarkIndexAny(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IndexAny("abcdefghijklmnopqrstuvwxyz", "zyxwvut")
	}
}
