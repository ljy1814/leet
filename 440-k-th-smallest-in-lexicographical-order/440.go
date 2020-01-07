package main

import "fmt"

func main() {
	n := 13
	k := 2
	n = 37
	k = 25
	n = 357
	k = 325
	fmt.Println(findKthNumber1(n, k))
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func getCount(prefix, n int) int {
	cnt := 0
	a := prefix
	b := prefix + 1

	for a <= n {
		cnt += min(n+1, b) - a
		a *= 10
		b *= 10
	}
	return cnt
}

func findKthNumber1(n int, k int) int {
	prefix := 1
	p := 1
	for p < k {
		cnt := getCount(prefix, n)
		if p+cnt > k {
			prefix *= 10
			p++
		} else if p+cnt <= k {
			prefix++
			p += cnt
		}
	}
	return prefix
}

func findKthNumber(n int, k int) int {
	if n < k || k < 1 {
		return 0
	}

	if n < 10 {
		return k
	}

	cur := 1
	k--

	for k > 0 {
		step := 0
		first := cur
		last := cur + 1
		for first <= n {
			step += min(n+1, last) - first
			first *= 10
			last *= 10
		}

		if step <= k {
			cur++
			k -= step
		} else {
			cur *= 10
			k--
		}
	}
	return cur
}
