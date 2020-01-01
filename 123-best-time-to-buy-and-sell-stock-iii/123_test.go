package main

import (
	"testing"
)

func TestOk(t *testing.T) {
	type (
		Tests struct {
			prices []int
			e      int
		}
	)
	ts := []*Tests{
		&Tests{
			prices: []int{3, 3, 5, 0, 0, 3, 1, 4},
			e:      6,
		},
		&Tests{
			prices: []int{1, 2, 3, 4, 5},
			e:      4,
		},
		&Tests{
			prices: []int{7, 6, 4, 3, 1},
			e:      0,
		},
	}
	for _, tt := range ts {
		g := maxProfit4(tt.prices)
		if g != tt.e {
			t.Fatalf("got:%d want:%d", g, tt.e)
		}
	}
}
