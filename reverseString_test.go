package challenges_tests

import "testing"

// Reverse returns the input string with its characters reversed
func Reverse(s string) string {

	r := ""
	for i := len(s) - 1; i >= 0; i-- {
		r = r + string(s[i])
	}

	return r
}

func TestReverse(t *testing.T) {

	var reverseTests = []struct {
		s   string
		out string
	}{
		{"asdf", "fdsa"},
		{"vU _!12 ?3P", "P3? 21!_ Uv"},
		{"", ""},
	}

	for _, test := range reverseTests {
		actual := Reverse(test.s)
		if actual != test.out {
			t.Errorf("Reverse(%q) = %q ; %q wanted", test.s, actual, test.out)
		}
	}

}

func BenchmarkReverse(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Reverse("abcdefghijklmnopqrstuvwxyz")
	}
}
