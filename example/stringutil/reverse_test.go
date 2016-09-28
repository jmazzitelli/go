package stringutil

import (
	"testing"
)

func TestReverseFunc(t *testing.T) {
	cases := []struct {
		input, expected string
	}{
		{"Hello World", "dlroW olleH"},
		{"", ""},
	}
	for _, c := range cases {
		got := Reverse(c.input)
		if got != c.expected {
			t.Errorf("Reverse(%s) == %s, expected %s", c.input, got, c.expected)
		}
	}
}
