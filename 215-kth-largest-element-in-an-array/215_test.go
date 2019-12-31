package main

import "testing"

func TestOk(t *testing.T) {
	type (
		Tests struct {
			nums []int
			k    int
			e    int
		}
	)
	ts := []*Tests{
		&Tests{
			nums: []int{
				3, 2, 1, 5, 6, 4,
			},
			k: 2,
			e: 5,
		},
		&Tests{
			nums: []int{
				3, 2, 3, 1, 2, 4, 5, 5, 6,
			},
			k: 4,
			e: 4,
		},
	}
	for _, tt := range ts {
		g := findKthLargest1(tt.nums, tt.k)
		if tt.e != g {
			t.Fatalf("got:%d expected:%d", g, tt.e)
		}
	}
}
