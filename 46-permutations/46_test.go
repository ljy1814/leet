package main

import (
	"fmt"
	"testing"
)

func TestOk(t *testing.T) {
	nums := []int{1, 2, 3}
	res := permute(nums)
	for _, r := range res {
		fmt.Printf("%p ", r)
		for _, a := range r {
			fmt.Print(a)
		}
		fmt.Println()
	}
}
