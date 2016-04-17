package lib

import (
	"fmt"
	"testing"
)

func TestRemoveLetters(t *testing.T) {
	type test struct {
		inS       string
		inExclude string
		outS      string
	}

	tests := []test{
		test{"abcde", "", "abcde"},
		test{"abcde", "abcde", ""},
		test{"abcdef", "bcdf", "ae"},
		test{"A_8!b.e?kd", "!_?", "A8b.ekd"},
	}

	for _, v := range tests {
		got := RemoveLetters(v.inS, v.inExclude)
		if got != v.outS {
			t.Error("fn(", v.inS, v.inExclude, ") =>", got, "!=", v.outS)
		} else {
			fmt.Println("PASS: fn(", v.inS, v.inExclude, ") =>", got, "==", v.outS)
		}
	}

}
