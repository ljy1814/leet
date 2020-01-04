package main

import "testing"

func TestOk(t *testing.T) {
	type (
		Tests struct {
			s string
			e int
		}
	)
	ts := []*Tests{
		&Tests{
			s: "(()",
			e: 2,
		},
		&Tests{
			s: ")()())",
			e: 4,
		},
	}

	for _, tt := range ts {
		g := longestValidParentheses(tt.s)
		if g != tt.e {
			t.Fatalf("got:%d want:%d", g, tt.e)
		}
	}
}
