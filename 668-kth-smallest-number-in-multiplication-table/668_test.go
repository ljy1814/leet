package main

import "testing"

func TestOk(t *testing.T) {
	type (
		Tests struct {
			m int
			n int
			k int
			e int
		}
	)

	ts := []*Tests{
		&Tests{
			m: 3,
			n: 3,
			k: 5,
			e: 3,
		},
	}

	for _, tt := range ts {
		g := findKthNumber(tt.m, tt.n, tt.k)
		if g != tt.e {
			t.Fatalf("got:%d want:%d", g, tt.e)
		}
	}
}
