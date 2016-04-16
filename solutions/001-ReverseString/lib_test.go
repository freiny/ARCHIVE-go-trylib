package lib

import "testing"

func TestReverse(t *testing.T) {
	var got, expect string
	got = Reverse("asdf")
	expect = "fdsa"
	if got != expect {
		t.Error("Expected ", expect, ", got ", got)
	}
}
