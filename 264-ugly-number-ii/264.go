package main

import "fmt"

func nthUglyNumber(n int) int {
	dp0 := []int{
		1, 2, 3,
	}

	if n < 3 {
		return dp0[n-1]
	}

	p2 := 1
	p3 := 0
	p5 := 0
	dp := make([]int, n)
	copy(dp, dp0)

	for i := 3; i < n; i++ {
		t2 := dp[p2] * 2
		for t2 <= dp[i-1] {
			p2++
			t2 = dp[p2] * 2
		}
		t3 := dp[p3] * 3
		for t3 <= dp[i-1] {
			p3++
			t3 = dp[p3] * 3
		}
		t5 := dp[p5] * 5
		for t5 <= dp[i-1] {
			p5++
			t5 = dp[p5] * 5
		}
		if t2 < t3 {
			if t2 > t5 {
				dp[i] = t5
			} else {
				dp[i] = t2
			}
		} else {
			if t3 > t5 {
				dp[i] = t5
			} else {
				dp[i] = t3
			}
		}
	}
	return dp[n-1]
}

func main() {
	n := 10
	fmt.Println(nthUglyNumber(n))
}
