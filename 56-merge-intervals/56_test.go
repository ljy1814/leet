package main

import (
	"fmt"
	"testing"
)

func TestOk(t *testing.T) {
	intervals := [][]int{
		[]int{1, 3}, []int{2, 6}, []int{8, 10}, []int{15, 18},
	}
	ret := merge(intervals)
	fmt.Println(ret)

	intervals = [][]int{
		[]int{1, 4}, []int{4, 5},
	}
	ret = merge(intervals)
	fmt.Println(ret)
}
