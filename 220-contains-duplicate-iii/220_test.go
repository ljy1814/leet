package main

import "testing"

func TestOk(t *testing.T) {
	type (
		Tests struct {
			nums []int
			k    int
			t    int
			e    bool
		}
	)

	ts := []*Tests{
		&Tests{
			nums: []int{
				1, 2, 3, 1,
			},
			k: 3,
			t: 0,
			e: true,
		},
		&Tests{
			nums: []int{
				1, 0, 1, 1,
			},
			k: 1,
			t: 2,
			e: true,
		},
		&Tests{
			nums: []int{
				1, 5, 9, 1, 5, 9,
			},
			k: 2,
			t: 3,
			e: false,
		},
		&Tests{
			nums: []int{
				-1, -1,
			},
			k: 1,
			t: 0,
			e: true,
		},
		&Tests{
			nums: []int{
				-3, 3,
			},
			k: 2,
			t: 4,
			e: false,
		},
	}

	for _, tt := range ts {
		g := containsNearbyAlmostDuplicate(tt.nums, tt.k, tt.t)
		if g != tt.e {
			t.Fatalf("got:%v want:%v nums:%+v", g, tt.e, tt.nums)
		}
	}
}
