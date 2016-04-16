package lib

// Reverse returns the input string with its characters reversed
func Reverse(s string) string {

	r := ""
	for i := len(s) - 1; i >= 0; i-- {
		r = r + string(s[i])
	}

	return r
}
