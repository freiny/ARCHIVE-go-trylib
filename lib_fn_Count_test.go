package fStrings

import (
	"strings"
	"testing"
)

func Count(s, sep string) int {
	if s == "" && sep == "" {
		return 1
	}
	if sep == "" {
		return len(s) + 1
	}
	count := 0
	sepLen := len(sep)
	slc := s[0:len(s)]
	for {
		switch {
		case len(slc) == sepLen:
			if slc == sep {
				count++
			}
			return count
		case len(slc) < sepLen:
			return count
		case len(slc) > sepLen:
			if slc[0:sepLen] == sep {
				count++
				slc = slc[sepLen:len(slc)]
			} else {
				slc = slc[1:len(slc)]
			}
		}
	}
}

var countTests = []struct {
	s, sep string
	out    int
}{
	{"", "", 1},
	{"a", "", 2},
	{"ab", "", 3},
	{"", "a", 0},
	{"a", "a", 1},
	{"a", "b", 0},
	{"a", "aaa", 0},
	{"abcde", "a", 1},
	{"abcde", "c", 1},
	{"abcde", "e", 1},
	{"abcde", "ab", 1},
	{"abcde", "bc", 1},
	{"abcde", "de", 1},
	{"aaa aaaa aaa", "aa", 4},
	{"aaa aaaa aaa", "aaaaa", 0},
}

func TestLibCount(t *testing.T) {
	for _, test := range countTests {
		actual := Count(test.s, test.sep)
		libout := strings.Count(test.s, test.sep)
		if actual != libout {
			t.Errorf("Count(%q, %q) = %d ; %d wanted", test.s, test.sep, actual, libout)
		}
	}
}

func TestCount(t *testing.T) {
	for _, test := range countTests {
		actual := Count(test.s, test.sep)
		if actual != test.out {
			t.Errorf("Count(%q, %q) = %d ; %d wanted", test.s, test.sep, actual, test.out)
		}
	}
}

func BenchmarkCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Count("aaba aaabab aaaa", "aa")
	}
}
