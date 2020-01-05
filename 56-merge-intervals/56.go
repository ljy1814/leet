package main

import "sort"

func main() {
}

func merge(intervals [][]int) [][]int {
	if len(intervals) < 1 || len(intervals[0]) < 2 {
		return nil
	}
	n := len(intervals)
	starts := make([]int, n+1)
	ends := make([]int, n)
	for i := 0; i < n; i++ {
		starts[i] = intervals[i][0]
		ends[i] = intervals[i][1]
	}
	starts[n] = 2<<31 - 1
	sort.Ints(starts)
	sort.Ints(ends)

	ret := [][]int{}
	// a b c d    b < c c>B -> -B] [c-
	// A B C D    B < C c<B -> ?->c->?->B] [a-b-c-B]
	for i, j := 0, 0; i < n; i++ {
		if starts[i+1] > ends[i] {
			ret = append(ret, []int{
				starts[j], ends[i],
			})
			j = i + 1
		}
	}

	return ret
}
