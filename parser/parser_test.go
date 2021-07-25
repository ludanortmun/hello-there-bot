package parser

import (
	"strconv"
	"testing"
)

func TestParser(t *testing.T) {
	cases := []struct {
		in   string
		want bool
	}{
		{"Hello, there", true},
		{"Something else", false},
		{"Hello there", true},
		{"Hello, there!", true},
		{"hello there", true},
		{"Hello  there", true},
		{"Hello, there 123", true},
		{" -heLLo, thErE", true},
	}

	for _, c := range cases {
		actual := Parse(c.in)
		if actual != c.want {
			t.Errorf("Parse(%q) == %q, expected %q", c.in, strconv.FormatBool(actual), strconv.FormatBool(c.want))
		}
	}
}
