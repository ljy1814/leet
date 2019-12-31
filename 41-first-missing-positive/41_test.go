package main

import "testing"

func TestOk(t *testing.T) {
	type (
		Tests struct {
			nums []int
			e    int
		}
	)
	ts := []*Tests{
		&Tests{
			nums: []int{
				1, 2, 0,
			},
			e: 3,
		},
		&Tests{
			nums: []int{
				3, 4, -1, 1,
			},
			e: 2,
		},
		&Tests{
			nums: []int{
				7, 8, 9, 11, 12,
			},
			e: 1,
		},
		&Tests{
			nums: []int{
				1, 1,
			},
			e: 2,
		},
	}

	for _, tt := range ts {
		g := firstMissingPositive(tt.nums)
		if g != tt.e {
			t.Fatalf("got:%d want:%d", g, tt.e)
		}
	}
}
